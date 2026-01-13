package render

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
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
		yaw:         0,
		pitch:       0,
		roll:        0,
		dislocation: Vec3{0, 0, 0},
		zoom:        100,
	}

	for !win.ShouldClose() {
		start := glfw.GetTime()
		delta := start - prev_time
		prev_time = start

		KeyboardInput(win, &param, delta)

		gl.ClearColor(1.0, 1.0, 1.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		x, y := win.GetSize()

		orientation := FromEuler(param.pitch, param.yaw, param.roll)
		roll_q := FromAxisAngle(z_vector, float32(Deg2Rad(float64(param.roll))))
		orientation = orientation.Mul(roll_q)

		rotation_mat := FromQuaternion(orientation)
		model := Translate(rotation_mat, param.dislocation)

		view := LookAt(camera_location, model.Mul4x1(Vec4{0, 0, 0, 1}).ToVec3(), world_up)

		shader.SetMatrix4f("model", model)

		proj := Perspective(float32(Deg2Rad(float64(param.zoom))), float32(x)/float32(y), 0.1, 1000.0)

		shader.SetMatrix4f("proj", proj)
		shader.SetMatrix4f("view", view)

		gl.DrawElements(gl.TRIANGLES, int32(len(vertices)), gl.UNSIGNED_INT, nil)
		win.SwapBuffers()
		glfw.PollEvents()
	}
}
