package render

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"math"
	. "scop/src/render/matrices"
	. "scop/src/render/transformations"
	. "scop/src/render/vectors"
	"scop/src/wavefront"
	"unsafe"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	FLOAT32_SIZE = int32(unsafe.Sizeof(float32(0)))
	UINT32_SIZE  = int32(unsafe.Sizeof(uint32(0)))
)

var (
	prev_time       = 0.0
	world_up        = Vec3{0, 1, 0}
	world_center    = Vec3{0, 0, 0}
	camera_location = Vec3{0, 0, 3}
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

	vbo := createVBO(1, [][]float32{wf.VerticesToVBO()})
	defer gl.DeleteBuffers(1, &vbo)

	// ebo := createEBO(1, [][]uint32{wf.FacesToEBO()})
	// defer gl.DeleteBuffers(1, &ebo)

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
		yaw:   0,
		pitch: 0,
		roll:  0,
		zoom:  100,
	}

	cube_pos := []Vec3{
		{0, 0, 0},
		{2, 5, -15},
		{-1.5, -2.2, -2.5},
		{-3.8, -2, -12.3},
		{2.4, -0.4, -3.5},
		{-1.7, 3, -7.5},
		{1.3, -2, -2.5},
		{1.5, 2, -2.5},
		{1.5, 0.2, -1.5},
		{-1.3, 1, -1.5},
	}

	for !win.ShouldClose() {
		start := glfw.GetTime()
		delta := start - prev_time
		prev_time = start

		KeyboardInput(win, &param, delta)

		gl.ClearColor(1.0, 1.0, 1.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		x, y := win.GetSize()
		view := LookAt(camera_location, world_center, world_up)
		proj := Perspective(float32(Deg2Rad(float64(param.zoom))), float32(x)/float32(y), 0.1, 100.0)

		shader.SetMatrix4f("proj", proj)
		shader.SetMatrix4f("view", view)

		pitch_sin, pitch_cos := math.Sincos(float64(param.pitch))
		yaw_sin, yaw_cos := math.Sincos(float64(param.yaw))
		front := Vec3{
			float32(pitch_cos * yaw_sin),
			float32(pitch_sin),
			float32(pitch_cos * yaw_cos),
		}.Normalize()

		for i := range 10 {
			right := world_up.CrossProduct(front).Normalize()
			up := front.CrossProduct(right).Normalize()

			model := Mat4{
				right.ToVec4(0),
				up.ToVec4(0),
				front.ToVec4(0),
				{0, 0, 0, 1},
			}

			model = Combine(
				model,
				Rotate(Identity4(), Deg2Rad(20*float64(i)), Vec3{1, 0.3, 0.5}),
				Translate(Identity4(), cube_pos[i]),
				Rotate(Identity4(), Deg2Rad(float64(param.roll)*40), Vec3{0, 0, 1}),
			)

			shader.SetMatrix4f("model", model)
			gl.DrawArrays(gl.TRIANGLES, 0, 36)
		}

		win.SwapBuffers()
		glfw.PollEvents()
	}
}
