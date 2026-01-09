package render

import (
	"image"
	"unsafe"

	"github.com/go-gl/gl/v4.5-core/gl"
)

func setViewport(x, y int) {
	gl.Viewport(0, 0, int32(x), int32(y))
}

func createVAO(n int32) uint32 {
	var vao uint32

	gl.GenVertexArrays(n, &vao)
	gl.BindVertexArray(vao)

	return vao
}

func createVBO(n int32, vertices [][]float32) uint32 {
	var vbo uint32

	gl.GenBuffers(n, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	for i := range vertices {
		gl.BufferData(gl.ARRAY_BUFFER, len(vertices[i])*int(FLOAT32_SIZE), gl.Ptr(&vertices[i][0]), gl.STATIC_DRAW)
	}

	return vbo
}

func createEBO(n int32, indexes [][]uint32) uint32 {
	var ebo uint32

	gl.GenBuffers(n, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	for i := range indexes {
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indexes[i])*int(UINT32_SIZE), gl.Ptr(&indexes[i][0]), gl.STATIC_DRAW)
	}

	return ebo
}

func setVertexAttibute(index uint32, size, stride int32, is_normalized bool, ptr unsafe.Pointer) {
	gl.VertexAttribPointer(index, size, gl.FLOAT, is_normalized, stride, ptr)
	gl.EnableVertexAttribArray(index)
}

func setVertexAttributeOffset(index uint32, size, stride int32, is_normalized bool, offset uintptr) {
	gl.VertexAttribPointerWithOffset(index, size, gl.FLOAT, is_normalized, stride, offset)
	gl.EnableVertexAttribArray(index)
}

func setTextureParameters() {
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
}

func createTexture(n int32, img *image.RGBA, texture_uint uint32) uint32 {
	var tex uint32

	width := int32(img.Rect.Dx())
	height := int32(img.Rect.Dy())

	gl.GenTextures(n, &tex)
	gl.ActiveTexture(texture_uint)
	gl.BindTexture(gl.TEXTURE_2D, tex)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&img.Pix[0]))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	return tex
}
