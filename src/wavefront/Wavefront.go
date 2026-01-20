package wavefront

import (
	"bufio"
	"fmt"
	"math"
	"os"
	fp "path/filepath"
	"scop/src"
	. "scop/src/render/vectors"
	mtlib "scop/src/wavefront/material_template_library"
	"strings"
)

type MTL = mtlib.MTL

type Wavefront struct {
	Mtl        *MTL
	Vertices   []*Vertex
	Material   *mtlib.Material
	Smoothing  uint8
	Normals    []*Normal
	Faces      []*Face
	Faces_type FaceType
}

func NewWavefront() *Wavefront {
	return &Wavefront{
		Mtl: mtlib.NewMTL(),
	}
}

func (wf *Wavefront) LoadData(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Printf("{INFO|WF} Start reading %s and parsing it\n", file.Name())

	for scanner.Scan() {
		line := src.ClearLineContent(scanner.Text())
		if line == "" {
			continue
		}

		tokens := strings.Fields(line)
		wf.matchAndUpdate(tokens)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	wf.Validate()

	fmt.Printf("{INFO|WF} Finish reading %s and parsing it\n", file.Name())
}

func (wf *Wavefront) matchAndUpdate(tokens []string) {
	switch strings.ToLower(tokens[0]) {
	case "mtllib":
		wf.addMTL(tokens)

	case "v":
		l := len(tokens)
		src.ValidateTokenAmount(l, 4, 5, "v must have three or four values")

		x, y, z := Str2Vertex(tokens[1:])

		var w float32 = 1.0
		var err error

		if l == 5 {
			w, err = src.ParseFloat[float32](tokens[4])
			if err != nil {
				panic(err)
			}
		}

		wf.Vertices = append(wf.Vertices, NewVertice(x, y, z, WithW(w)))

	case "usemtl":
		if wf.Mtl == nil {
			panic("mtllib not defined")
		}
		src.ValidateTokenAmount(len(tokens), 2, 2, "usemtl must have one value")
		wf.selectMaterial(tokens[1])

	case "s":
		src.ValidateTokenAmount(len(tokens), 2, 2, "s must have one value")
		if tokens[1] == "off" {
			wf.Smoothing = 0
			return
		}

		n, err := src.ParseUint[uint8](tokens[1])
		if err != nil {
			panic(err)
		}

		wf.Smoothing = n

	case "f":
		src.ValidateTokenAmount(len(tokens), 4, math.MaxInt, "f must have at least four values")

		faces_len := len(wf.Faces)

		for i, tok := range tokens[1:] {
			values := strings.Split(tok, "/")
			if len(values) > 3 {
				panic("face internal value has more than three components")
			}

			if wf.Faces_type == NOT_SET {
				wf.setFacesType(values)
			} else {
				wf.validateFacesType(values)
			}

			v, t, n := Tokens2FaceSlices(values)
			if i == 0 {
				wf.Faces = append(wf.Faces, NewFace(wf.Faces_type, v, WithTexture(t), WithNormals(n)))
			} else {
				face := wf.Faces[faces_len]
				face.AppendVertex(v)
				face.AppendTexture(t)
				face.AppendNormal(n)
			}
		}

	case "vn":
		src.ValidateTokenAmount(len(tokens), 4, 4, "vn must have only three values")

		x, err := src.ParseFloat[float32](tokens[1])
		if err != nil {
			panic(err)
		}
		y, err := src.ParseFloat[float32](tokens[2])
		if err != nil {
			panic(err)
		}
		z, err := src.ParseFloat[float32](tokens[3])
		if err != nil {
			panic(err)
		}

		wf.Normals = append(wf.Normals, NewNormal(x, y, z))

	case "o", "g":
		return

	default:
		panic(fmt.Sprintf("Unknown identifier '%s'", tokens[0]))
	}
}

func (wf *Wavefront) addMTL(tokens []string) {
	filename := "/" + strings.Join(tokens[1:], " ")
	filename = fp.Dir(os.Args[1]) + filename

	wf.Mtl = mtlib.NewMTL()
	wf.Mtl.LoadData(filename)
}

func (wf *Wavefront) selectMaterial(name string) {
	for _, mat := range wf.Mtl.Materials {
		if mat.Name == name {
			wf.Material = mat
			break
		}
	}

	if wf.Material == nil && name != "Default" {
		panic(fmt.Sprintf("material of name '%s', not found in loaded .mtl", name))
	} else if name == "Default" {
		wf.Material = wf.Mtl.DefaultMaterial()
	}
}

func (wf *Wavefront) Validate() {
	if wf.Mtl == nil {
		panic("mtllib not defined")
	}

	if wf.Smoothing != 0 && wf.Normals == nil {
		panic("can't define smoothing without normals")
	}
}

func (wf *Wavefront) setFacesType(values []string) {
	switch len(values) {
	case 1:
		wf.Faces_type = VERTICE
	case 2:
		wf.Faces_type = VERTICE_TEXTURE
	case 3:
		if values[1] != "" {
			wf.Faces_type = VERTICE_TEXTURE_NORMAL
		} else {
			wf.Faces_type = VERTICE_NORMAL
		}
	}
}

func (wf *Wavefront) validateFacesType(values []string) {
	switch len(values) {
	case 1:
		if wf.Faces_type != VERTICE {
			panic("face type mismatch")
		}
	case 2:
		if wf.Faces_type != VERTICE_TEXTURE {
			panic("face type mismatch")
		}
	case 3:
		if values[1] != "" {
			if wf.Faces_type != VERTICE_TEXTURE_NORMAL {
				panic("face type mismatch")
			}
		} else {
			if wf.Faces_type != VERTICE_NORMAL {
				panic("face type mismatch")
			}
		}
	}
}

func Tokens2FaceSlices(tokens []string) (vertice int, texture int, normal int) {
	var err error

	vertice, err = src.ParseInt[int](tokens[0])
	if err != nil {
		panic(err)
	} else if vertice == 0 {
		panic("face vertice index cannot be zero")
	}

	if len(tokens) > 1 && tokens[1] != "" {
		texture, err = src.ParseInt[int](tokens[1])
		if err != nil {
			panic(err)
		} else if texture == 0 {
			panic("face texture index cannot be zero")
		}
	}

	if len(tokens) > 2 {
		normal, err = src.ParseInt[int](tokens[2])
		if err != nil {
			panic(err)
		} else if normal == 0 {
			panic("face normal index cannot be zero")
		}
	}

	return
}

func (wf *Wavefront) VerticesToVBO() []float32 {
	var positions []float32

	for _, v := range wf.Vertices {
		positions = append(positions, v.X, v.Y, v.Z, v.W)
	}

	min, max := computeBounds(wf.Vertices)
	width := max[0] - min[0]
	height := max[1] - min[1]
	depth := max[2] - min[2]
	if width == 0 {
		width = 1.0
	}
	if height == 0 {
		height = 1.0
	}
	if depth == 0 {
		depth = 1.0
	}

	var vertices []float32
	for i := 0; i < len(positions); i += 4 {
		x, y, z, w := positions[i], positions[i+1], positions[i+2], positions[i+3]
		u := (z - min[2]) / depth * 3
		v := (y - min[1]) / height * 2
		vertices = append(vertices, x, y, z, w, u, v)
	}

	return vertices
}

func (wf *Wavefront) FacesToEBO() []uint32 {
	var indices []uint32

	for _, face := range wf.Faces {
		if len(face.Vertices_indexes) > 3 {
			indices = append(indices, uint32(face.Vertices_indexes[0]-1), uint32(face.Vertices_indexes[1]-1), uint32(face.Vertices_indexes[2]-1))
			indices = append(indices, uint32(face.Vertices_indexes[0]-1), uint32(face.Vertices_indexes[2]-1), uint32(face.Vertices_indexes[3]-1))
		} else {
			for _, index := range face.Vertices_indexes {
				indices = append(indices, uint32(index-1))
			}
		}
	}

	return indices
}

func computeBounds(vertices []*Vertex) (min Vec3, max Vec3) {
	min = Vec3{math.MaxFloat32, math.MaxFloat32, math.MaxFloat32}
	max = Vec3{-math.MaxFloat32, -math.MaxFloat32, -math.MaxFloat32}

	for _, v := range vertices {
		min[0] = float32(math.Min(float64(min[0]), float64(v.X)))
		min[1] = float32(math.Min(float64(min[1]), float64(v.Y)))
		min[2] = float32(math.Min(float64(min[2]), float64(v.Z)))

		max[0] = float32(math.Max(float64(max[0]), float64(v.X)))
		max[1] = float32(math.Max(float64(max[1]), float64(v.Y)))
		max[2] = float32(math.Max(float64(max[2]), float64(v.Z)))
	}

	return min, max
}
