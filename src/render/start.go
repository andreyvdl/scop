package render

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
	"scop/src/render/matrices"
	"scop/src/render/transformations"
	"scop/src/render/vectors"
	"scop/src/wavefront"
	"time"
	"unsafe"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	FLOAT32_SIZE = int32(unsafe.Sizeof(float32(0)))
	UINT32_SIZE  = int32(unsafe.Sizeof(uint32(0)))
	FRAME_TIME   = time.Second / time.Duration(60)
)

func StartRendering(wf *wavefront.Wavefront) {
	_ = wf

	fmt.Println("{INFO|START} Starting window render")
	defer fmt.Println("{INFO|START} Stopping window render")

	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	win := createWindow()
	defer win.Destroy()
	win.MakeContextCurrent()

	// glfw.SwapInterval(1)

	fmt.Println("{INFO|START} Starting opengl")
	defer fmt.Println("{INFO|START} Stopping opengl")

	err = gl.Init()
	if err != nil {
		panic(err)
	}

	var nr_attributes int32
	gl.GetIntegerv(gl.MAX_VERTEX_ATTRIBS, &nr_attributes)
	fmt.Println("Max number of vertex attributes:", nr_attributes)

	w, h := win.GetSize()
	gl.Viewport(0, 0, int32(w), int32(h))
	win.SetFramebufferSizeCallback(FBSCallback)

	vertices := wf.VerticesToVBO()
	indexes := wf.FacesToEBO()

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*int(FLOAT32_SIZE), gl.Ptr(&vertices[0]), gl.STATIC_DRAW)

	var ebo uint32
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indexes)*int(UINT32_SIZE), gl.Ptr(&indexes[0]), gl.STATIC_DRAW)

	shader := NewShader("src/render/vertex.vs.glsl", "src/render/fragment.fs.glsl")
	shader.Use()

	gl.VertexAttribPointer(0, 4, gl.FLOAT, false, 6*FLOAT32_SIZE, nil) // vertex
	gl.EnableVertexAttribArray(0)                                      // need to call this before the color

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.MIRRORED_REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.MIRRORED_REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST_MIPMAP_NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	img1 := loadTexture("resources/goldo_shippu.jpg")
	img2 := loadTexture("resources/sombra.png")
	flipVertically(img1)
	flipVertically(img2)

	var tex1 uint32
	var tex2 uint32
	gl.GenTextures(1, &tex1)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, tex1)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(img1.Rect.Dx()), int32(img1.Rect.Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&img1.Pix[0]))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	gl.GenTextures(1, &tex2)
	gl.ActiveTexture(gl.TEXTURE1)
	gl.BindTexture(gl.TEXTURE_2D, tex2)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(img2.Rect.Dx()), int32(img2.Rect.Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&img2.Pix[0]))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	shader.SetInt("tex1", 0)
	shader.SetInt("tex2", 1)

	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 6*FLOAT32_SIZE, uintptr(4*FLOAT32_SIZE)) // texture coordinates
	gl.EnableVertexAttribArray(1)

	for !win.ShouldClose() {
		start := time.Now()
		m := matrices.Identity4()
		m = transformations.Combine(m,
			transformations.Rotate(m, Deg2Rad(90), vectors.Vec3{0, 0, 1}),
			transformations.Scale(m, vectors.Vec3{0.5, 0.5, 0.5}),
		)

		KeyboardInput(win)
		shader.SetMatrix4f("transform", m)

		gl.ClearColor(0.0, 0.0, 0.314, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		shaderColor(shader)

		gl.BindTexture(gl.TEXTURE_2D, tex1)
		gl.BindTexture(gl.TEXTURE_2D, tex2)
		gl.BindVertexArray(vao)
		gl.DrawElements(gl.TRIANGLES, int32(len(indexes)), gl.UNSIGNED_INT, nil)

		win.SwapBuffers()
		glfw.PollEvents()

		elapsed := time.Since(start)
		if sleep_duration := FRAME_TIME - elapsed; sleep_duration >= 0 {
			time.Sleep(sleep_duration)
		} else {
			fmt.Println("{WARNING} FPS TOO LOW")
		}
	}
}

func createWindow() *glfw.Window {
	video := glfw.GetPrimaryMonitor().GetVideoMode()
	width := video.Width / 2
	height := video.Height / 2

	win, err := glfw.CreateWindow(width, height, "scop", nil, nil)
	if err != nil {
		panic(err)
	}

	return win
}

func shaderColor(shader *Shader) {
	timeValue := glfw.GetTime()
	green := math.Tan(timeValue)/2.0 + 0.5

	shader.Use()
	shader.Set4Float("ourColor", float32(math.Pow(green, 2)), float32(green), float32(green)/2, 1.0)
}

func loadTexture(filepath string) *image.RGBA {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, img.Bounds().Min, draw.Src)

	return rgba
}

func flipVertically(img *image.RGBA) {
	row := make([]byte, img.Bounds().Dx()*4)
	for y := 0; y < img.Bounds().Dy()/2; y++ {
		top := y * img.Stride
		bottom := (img.Bounds().Dy() - 1 - y) * img.Stride
		copy(row, img.Pix[top:top+img.Stride])
		copy(img.Pix[top:top+img.Stride], img.Pix[bottom:bottom+img.Stride])
		copy(img.Pix[bottom:bottom+img.Stride], row)
	}

}
