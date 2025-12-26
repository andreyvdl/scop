package wavefront

type FaceType uint8

const (
	NOT_SET FaceType = iota
	VERTICE
	VERTICE_TEXTURE
	VERTICE_TEXTURE_NORMAL
	VERTICE_NORMAL
)

type Face struct {
	Type             FaceType
	Vertices_indexes []int
	Texture_indexes  []int
	Normals_indexes  []int
}

type faceOptions func(*Face)

func NewFace(typ FaceType, vertex int, opts ...faceOptions) *Face {
	f := &Face{Type: typ}

	f.Vertices_indexes = append(f.Vertices_indexes, vertex)

	for _, opt := range opts {
		opt(f)
	}

	return f
}

func (f *Face) AppendVertex(vertices int) {
	f.Vertices_indexes = append(f.Vertices_indexes, vertices)
}

func (f *Face) AppendTexture(texture int) {
	f.Texture_indexes = append(f.Texture_indexes, texture)
}

func (f *Face) AppendNormal(normal int) {
	f.Normals_indexes = append(f.Normals_indexes, normal)
}

func WithTexture(textureIndex int) faceOptions {
	return func(f *Face) {
		f.Texture_indexes = append(f.Texture_indexes, textureIndex)
	}
}

func WithNormals(normalsIndex int) faceOptions {
	return func(f *Face) {
		f.Normals_indexes = append(f.Normals_indexes, normalsIndex)
	}
}
