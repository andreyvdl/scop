package materialtemplatelibrary

import (
	"bufio"
	"fmt"
	"os"
	"scop/src"
	"strings"
)

type MTL struct {
	Materials []*Material
}

func NewMTL() *MTL {
	return &MTL{}
}

func (m *MTL) LoadData(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	size := 0

	fmt.Printf("{INFO|MTL} Start reading %s and parsing it\n", file.Name())

	for scanner.Scan() {
		line := src.ClearLineContent(scanner.Text())
		if line == "" {
			continue
		}

		tokens := strings.Fields(line)
		m.matchAndUpdate(tokens, &size)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	checkNotFilledMaterials(m)
	fmt.Printf("{INFO|MTL} Finish reading %s and parsing it\n", file.Name())
}

func (m *MTL) matchAndUpdate(tokens []string, size *int) {
	switch strings.ToLower(tokens[0]) {
	case "newmtl":
		src.ValidateTokenAmount(len(tokens), 2, 2, "newmtl must have one value")
		checkDuplicateName(m.Materials, tokens[1])
		m.Materials = append(m.Materials, NewMaterial(tokens[1]))
		(*size)++

	case "ns":
		checkMaterialDeclared(*size)
		src.ValidateTokenAmount(len(tokens), 2, 2, "ns must have one value")

		n, err := src.ParseFloat[float32](tokens[1])
		if err != nil {
			panic(err)
		}
		m.Materials[(*size)-1].Specular_exponent = NewNs(n)

	case "ka":
		checkMaterialDeclared(*size)
		src.ValidateTokenAmount(len(tokens), 4, 4, "ka must have three values")

		r, g, b := src.Str2RGB(tokens[1:])
		m.Materials[(*size)-1].Ambient_color = NewKa(r, g, b)

	case "kd":
		checkMaterialDeclared(*size)
		src.ValidateTokenAmount(len(tokens), 4, 4, "kd must have three values")

		r, g, b := src.Str2RGB(tokens[1:])
		m.Materials[(*size)-1].Diffuse_color = NewKd(r, g, b)

	case "ks":
		checkMaterialDeclared(*size)
		src.ValidateTokenAmount(len(tokens), 4, 4, "ks must have three values")

		r, g, b := src.Str2RGB(tokens[1:])
		m.Materials[(*size)-1].Specular_color = NewKs(r, g, b)

	case "ni":
		checkMaterialDeclared(*size)
		src.ValidateTokenAmount(len(tokens), 2, 2, "ni must have one value")

		n, err := src.ParseFloat[float32](tokens[1])
		if err != nil {
			panic(err)
		}
		m.Materials[(*size)-1].Refraction_coefficient = NewNi(n)

	case "d":
		checkMaterialDeclared(*size)
		l := len(tokens)
		src.ValidateTokenAmount(l, 2, 3, "d must have one or two values")

		f, err := src.ParseFloat[float32](tokens[l-1])
		if err != nil {
			panic(err)
		}

		style := DEFAULT_STYLE
		if l == 3 {
			if strings.ToLower(tokens[1]) != "-halo" {
				panic(fmt.Sprintf("Invalid dissolve style: '%s'", tokens[1]))
			}
			style = HALO_STYLE
		}
		m.Materials[(*size)-1].Dissolve_factor = NewD(f, WithStyle(style))

	case "illum":
		checkMaterialDeclared(*size)
		src.ValidateTokenAmount(len(tokens), 2, 2, "illum must have one value")
		m.Materials[(*size)-1].Illumination_model = NewIllum(tokens[1])

	default:
		panic(fmt.Sprintf("Unknown .mtl identifier: '%s'", tokens[0]))
	}
}

func checkMaterialDeclared(size int) {
	if size == 0 {
		panic("No material declared before assigning values")
	}
}

func checkNotFilledMaterials(m *MTL) {
	for _, mat := range m.Materials {
		if mat.Ambient_color == nil {
			panic("Ambient color (Ka) not defined")
		} else if mat.Diffuse_color == nil {
			panic("Diffuse color (Kd) not defined")
		} else if mat.Illumination_model == nil {
			panic("Illumination model (illum) not defined")
		} else if mat.Refraction_coefficient == nil {
			panic("Refraction coefficient (Ni) not defined")
		} else if mat.Specular_color == nil {
			panic("Specular color (Ks) not defined")
		} else if mat.Specular_exponent == nil {
			panic("Specular exponent (Ns) not defined")
		} else if mat.Dissolve_factor == nil {
			fmt.Println("{WARNING|MTL} dissolve factor (d) not defined. Assuming default values")
			mat.Dissolve_factor = &D{float32(1.0), DEFAULT_STYLE}
		}
	}
}

func checkDuplicateName(materials []*Material, name string) {
	for _, mat := range materials {
		if mat.Name == name {
			panic(fmt.Sprintf("Duplicate material name: '%s'", name))
		}
	}
}

func (m *MTL) DefaultMaterial() *Material {
	m.Materials = append(
		m.Materials,
		&Material{
			Name:                   "Default",
			Ambient_color:          NewKa(1.0, 1.0, 1.0),
			Diffuse_color:          NewKd(1.0, 1.0, 1.0),
			Specular_color:         NewKs(0.0, 0.0, 0.0),
			Refraction_coefficient: NewNi(1.0),
			Specular_exponent:      NewNs(0.0),
			Dissolve_factor:        NewD(1.0),
			Illumination_model:     NewIllum("0"),
		},
	)
	return m.Materials[len(m.Materials)-1]
}
