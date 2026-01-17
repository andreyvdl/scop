package render

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"

	. "scop/src/render/vectors"
	"scop/src/wavefront"
	"time"
	"unsafe"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	FLOAT32_SIZE = int32(unsafe.Sizeof(float32(0)))
	UINT32_SIZE  = int32(unsafe.Sizeof(uint32(0)))
)

var (
	prev_time       = time.Now()
	world_up        = Vec3{0, 1, 0}
	world_center    = Vec3{0, 0, 0}
	camera_location = Vec3{0, 0, -5}
	z_vector        = Vec3{0, 0, 1}
)

func StartRendering(wf *wavefront.Wavefront) {
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

	setViewport(win.GetSize())
	win.SetFramebufferSizeCallback(FBSCallback)

	vao := createVAO(1)
	defer gl.DeleteVertexArrays(1, &vao)

	vertices := wf.VerticesToVBO()
	vbo := createVBO(1, [][]float32{vertices})
	defer gl.DeleteBuffers(1, &vbo)

	ebo := createEBO(1, [][]uint32{wf.FacesToEBO()})
	defer gl.DeleteBuffers(1, &ebo)

	shader := NewShader("src/render/vertex.vs.glsl", "src/render/fragment.fs.glsl")
	shader.Use()

	setVertexAttibute(0, 4, 6*FLOAT32_SIZE, false, nil) // vertices
	setTextureParameters()

	img1 := loadTexture("resources/goldo_shippu.jpg")
	img2 := loadTexture("resources/sombra.png")
	flipVertically(img1)
	flipVertically(img2)
	createTexture(1, img1, gl.TEXTURE0)
	createTexture(1, img2, gl.TEXTURE1)
	shader.SetInt("tex1", 0)
	shader.SetInt("tex2", 1)

	setVertexAttributeOffset(1, 2, 6*FLOAT32_SIZE, false, uintptr(4*FLOAT32_SIZE)) // texture coordinates
	gl.Enable(gl.DEPTH_TEST)

	param := Parameters{
		yaw:             0,
		pitch:           0,
		roll:            0,
		dislocation:     Vec3{0, 0, 0},
		zoom:            100,
		fps:             60,
		transition_side: 0.01,
		transition:      1.0,
	}

	loop(win, shader, &param, int32(len(vertices)))
}

func loop(win *glfw.Window, shader *Shader, param *Parameters, vertices_len int32) {
	for !win.ShouldClose() {
		frame_time := time.Second / time.Duration(param.fps)
		start := time.Now()
		delta := start.Sub(prev_time).Seconds()
		prev_time = start

		KeyboardInput(win, param, delta)

		gl.ClearColor(1.0, 1.0, 1.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		shader.SetFloat("transition", param.transition)

		sendTransformationDataToShader(shader, param, win)

		gl.DrawElements(gl.TRIANGLES, vertices_len, gl.UNSIGNED_INT, nil)
		win.SwapBuffers()
		glfw.PollEvents()

		param.transition += param.transition_side
		if param.transition < 0.0 {
			param.transition = 0.0
		} else if param.transition > 1.0 {
			param.transition = 1.0
		}

		if rest := frame_time - time.Since(start); rest > 0 {
			time.Sleep(rest)
		}
	}
}

func sendTransformationDataToShader(shader *Shader, param *Parameters, win *glfw.Window) {
	shader.SetFloat("transition", param.transition)

	shader.SetFloat("pitch", param.pitch)
	shader.SetFloat("yaw", param.yaw)
	shader.SetFloat("roll", param.roll)
	shader.Set3Float("z_vec", z_vector[0], z_vector[1], z_vector[2])

	shader.Set3Float("dislocation", param.dislocation[0], param.dislocation[1], param.dislocation[2])
	shader.Set3Float("camera", camera_location[0], camera_location[1], camera_location[2])
	shader.Set3Float("world_up", world_up[0], world_up[1], world_up[2])

	x, y := win.GetSize()
	shader.SetFloat("width", float32(x))
	shader.SetFloat("height", float32(y))
	shader.SetFloat("zoom", param.zoom)
}
