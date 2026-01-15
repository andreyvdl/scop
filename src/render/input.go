package render

import (
	"math"
	. "scop/src/render/vectors"

	"github.com/go-gl/glfw/v3.3/glfw"
)

var last_call = 0.0

func KeyboardInput(win *glfw.Window, param *Parameters, delta float64) {
	if win.GetKey(glfw.KeyEscape) == glfw.Press {
		win.SetShouldClose(true)
		return
	}

	delta32 := float32(delta)

	checkReset(win, param)
	checkVSync(win)
	checkRotations(win, param, delta32)
	checkDislocation(win, param, delta32)
	checkZoom(win, param, delta32)
	checkFPS(win, param)
	checkTexture(win, param)
}

func checkReset(win *glfw.Window, param *Parameters) {
	if win.GetKey(glfw.KeyR) == glfw.Press {
		param.yaw = 0
		param.pitch = 0
		param.roll = 0
		param.zoom = 100
		param.dislocation = Vec3{0, 0, 0}
	}
}

func checkVSync(win *glfw.Window) {
	if win.GetKey(glfw.KeyV) == glfw.Press && win.GetKey(glfw.KeyLeftShift) == glfw.Press {
		glfw.SwapInterval(0)
	} else if win.GetKey(glfw.KeyV) == glfw.Press {
		glfw.SwapInterval(1)
	}
}

func checkRotations(win *glfw.Window, param *Parameters, delta float32) {
	speed := 100 * delta

	if win.GetKey(glfw.KeyUp) == glfw.Press {
		param.pitch -= speed
	}
	if win.GetKey(glfw.KeyDown) == glfw.Press {
		param.pitch += speed
	}

	if win.GetKey(glfw.KeyLeft) == glfw.Press {
		param.yaw += speed
	}
	if win.GetKey(glfw.KeyRight) == glfw.Press {
		param.yaw -= speed
	}

	if win.GetKey(glfw.KeyQ) == glfw.Press {
		param.roll -= speed
	}
	if win.GetKey(glfw.KeyE) == glfw.Press {
		param.roll += speed
	}

	param.pitch = float32(math.Mod(float64(param.pitch), 360))
	param.yaw = float32(math.Mod(float64(param.yaw), 360))
	param.roll = float32(math.Mod(float64(param.roll), 360))
}

func checkDislocation(win *glfw.Window, param *Parameters, delta float32) {
	speed := 10 * delta

	if win.GetKey(glfw.KeyW) == glfw.Press {
		param.dislocation[1] += speed
	}
	if win.GetKey(glfw.KeyS) == glfw.Press {
		param.dislocation[1] -= speed
	}

	if win.GetKey(glfw.KeyA) == glfw.Press {
		param.dislocation[0] -= speed
	}
	if win.GetKey(glfw.KeyD) == glfw.Press {
		param.dislocation[0] += speed
	}

	if win.GetKey(glfw.KeyN) == glfw.Press {
		param.dislocation[2] -= speed
	}
	if win.GetKey(glfw.KeyM) == glfw.Press {
		param.dislocation[2] += speed
	}
}

func checkZoom(win *glfw.Window, param *Parameters, delta float32) {
	scale := 24 * delta

	if win.GetKey(glfw.Key1) == glfw.Press && param.zoom > 0 {
		param.zoom -= scale
	}
	if win.GetKey(glfw.Key3) == glfw.Press && param.zoom < 180 {
		param.zoom += scale
	}
}

func checkFPS(win *glfw.Window, param *Parameters) {
	if win.GetKey(glfw.KeyMinus) == glfw.Press {
		param.fps -= 1
	}
	if win.GetKey(glfw.KeyEqual) == glfw.Press {
		param.fps += 1
	}

	if param.fps < 30 {
		param.fps = 30
	} else if param.fps > 260 {
		param.fps = 260
	}
}

func checkTexture(win *glfw.Window, param *Parameters) {
	if win.GetKey(glfw.KeyT) == glfw.Press {
		if glfw.GetTime()-last_call < 0.2 { // avoid multiple calls in the same frame
			return
		}

		param.transition_side = -param.transition_side

		last_call = glfw.GetTime()
	}
}
