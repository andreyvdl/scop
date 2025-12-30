package render

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
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

	gl.VertexAttribPointer(0, 4, gl.FLOAT, false, 9*FLOAT32_SIZE, nil) // vertex
	gl.EnableVertexAttribArray(0)                                      // need to call this before the color

	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 9*FLOAT32_SIZE, uintptr(4*FLOAT32_SIZE)) // color
	gl.EnableVertexAttribArray(1)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.MIRRORED_REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.MIRRORED_REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST_MIPMAP_NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	gs, err := os.Open("resources/goldo_shippu.jpg")
	// gs, err := os.Open("resources/sombra.png")
	if err != nil {
		panic(err)
	}
	defer gs.Close()

	img, _, err := image.Decode(gs)
	if err != nil {
		panic(err)
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, img.Bounds().Min, draw.Src)

	var tex uint32
	gl.GenTextures(1, &tex)
	gl.BindTexture(gl.TEXTURE_2D, tex)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(rgba.Rect.Dx()), int32(rgba.Rect.Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&rgba.Pix[0]))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	gl.VertexAttribPointerWithOffset(2, 2, gl.FLOAT, false, 9*FLOAT32_SIZE, uintptr(7*FLOAT32_SIZE)) // texture coordinates
	gl.EnableVertexAttribArray(2)

	for !win.ShouldClose() {
		start := time.Now()

		KeyboardInput(win)

		gl.ClearColor(0.0, 0.0, 0.314, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		shaderColor(shader)

		gl.BindTexture(gl.TEXTURE_2D, tex)
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
