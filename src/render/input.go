package render

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

func KeyboardInput(win *glfw.Window, param *Parameters, delta float64) {
	if win.GetKey(glfw.KeyEscape) == glfw.Press {
		win.SetShouldClose(true)
	}

	if win.GetKey(glfw.KeyUp) == glfw.Press {
		param.rot[0] += 100 * delta
		if param.rot[0] >= 360 {
			param.rot[0] = 0
		}
	}
	if win.GetKey(glfw.KeyDown) == glfw.Press {
		param.rot[0] -= 100 * delta
		if param.rot[0] <= 0 {
			param.rot[0] = 359
		}
	}

	if win.GetKey(glfw.KeyLeft) == glfw.Press {
		param.rot[1] += 100 * delta
		if param.rot[1] >= 360 {
			param.rot[1] = 0
		}
	}
	if win.GetKey(glfw.KeyRight) == glfw.Press {
		param.rot[1] -= 100 * delta
		if param.rot[1] <= 0 {
			param.rot[1] = 359
		}
	}

	if win.GetKey(glfw.KeyX) == glfw.Press {
		param.rot[2] += 100 * delta
		if param.rot[2] >= 360 {
			param.rot[2] = 0
		}
	}
	if win.GetKey(glfw.KeyZ) == glfw.Press {
		param.rot[2] -= 100 * delta
		if param.rot[2] <= 0 {
			param.rot[2] = 359
		}
	}

	if win.GetKey(glfw.Key1) == glfw.Press && param.scale > 0.01 {
		param.scale -= 1 * float32(delta)
	}
	if win.GetKey(glfw.Key3) == glfw.Press && param.scale < 1.99 {
		param.scale += 1 * float32(delta)
	}
}
