package render

import (
	_ "embed"
	"fmt"
	"scop/src/wavefront"
	"unsafe"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

//go:embed vertex.glsl
var vertex_source []byte

//go:embed fragment.glsl
var fragment_source []byte

const FLOAT32_SIZE = int32(unsafe.Sizeof(float32(0)))

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

	glfw.SwapInterval(1)

	fmt.Println("{INFO|START} Starting opengl")
	defer fmt.Println("{INFO|START} Stopping opengl")

	err = gl.Init()
	if err != nil {
		panic(err)
	}

	w, h := win.GetSize()
	gl.Viewport(0, 0, int32(w), int32(h))
	win.SetFramebufferSizeCallback(FBSCallback)

	vertices := []float32{
		-0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, 0.0, 1.0,
		0.0, 0.5, 0.0, 1.0,
		0.0, -0.5, 0.0, 1.0,
	}

	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*int(FLOAT32_SIZE), gl.Ptr(&vertices[0]), gl.STATIC_DRAW)

	initShaders()

	gl.VertexAttribPointer(0, 4, gl.FLOAT, false, 4*FLOAT32_SIZE, nil)
	gl.EnableVertexAttribArray(0)

	gl.BindVertexArray(VAO)

	for !win.ShouldClose() {
		KeyboardInput(win)
		gl.ClearColor(0.0, 0.0, 0.314, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/4))
		win.SwapBuffers()
		glfw.PollEvents()
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

func initShaders() {
	vtSrc := string(vertex_source)
	vtSrc += "\x00"
	vt_uint, free := gl.Strs(vtSrc)
	defer free()
	vertex_shader := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(vertex_shader, 1, vt_uint, nil)
	gl.CompileShader(vertex_shader)

	status := int32(0)
	gl.GetShaderiv(vertex_shader, gl.COMPILE_STATUS, &status)
	if status == 0 {
		log := make([]byte, 1024)
		gl.GetShaderInfoLog(vertex_shader, 1024, nil, &log[0])
		panic(fmt.Errorf("Vertex shader compilation error: %s", string(log)))
	} else {
		fmt.Println("{INFO|START} Vertex shader compiled")
	}
	defer gl.DeleteShader(vertex_shader)

	fgSrc := string(fragment_source)
	fgSrc += "\x00"
	fg_uint, free := gl.Strs(fgSrc)
	defer free()
	fragment_shader := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(fragment_shader, 1, fg_uint, nil)
	gl.CompileShader(fragment_shader)

	status = int32(0)
	gl.GetShaderiv(fragment_shader, gl.COMPILE_STATUS, &status)
	if status == 0 {
		log := make([]byte, 1024)
		gl.GetShaderInfoLog(fragment_shader, 1024, nil, &log[0])
		panic(fmt.Errorf("Fragment shader compilation error: %s", string(log)))
	} else {
		fmt.Println("{INFO|START} Fragment shader compiled")
	}
	defer gl.DeleteShader(fragment_shader)

	program := gl.CreateProgram()
	gl.AttachShader(program, vertex_shader)
	gl.AttachShader(program, fragment_shader)
	gl.LinkProgram(program)

	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == 0 {
		log := make([]byte, 1024)
		gl.GetProgramInfoLog(program, 1024, nil, &log[0])
		panic(fmt.Errorf("Program linking error: %s", string(log)))
	} else {
		fmt.Println("{INFO|START} Program linked")
	}
	gl.UseProgram(program)
}
