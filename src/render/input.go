package render

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

func KeyboardInput(win *glfw.Window, param *Parameters, delta float64) {
	if win.GetKey(glfw.KeyEscape) == glfw.Press {
		win.SetShouldClose(true)
		return
	}

	checkReset(win, param)
	checkVSync(win)
	checkRotations(win, param, float32(delta))
	checkZoom(win, param, float32(delta))
}

func checkReset(win *glfw.Window, param *Parameters) {
	if win.GetKey(glfw.KeyR) == glfw.Press {
		param.yaw = 0
		param.pitch = 0
		param.roll = 0
		param.zoom = 100
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
	speed := 3 * delta

	if win.GetKey(glfw.KeyUp) == glfw.Press {
		param.pitch -= speed
		if param.pitch < 0 {
			param.pitch += 360
		}
	}
	if win.GetKey(glfw.KeyDown) == glfw.Press {
		param.pitch += speed
		if param.pitch > 360 {
			param.pitch -= 360
		}
	}

	if win.GetKey(glfw.KeyLeft) == glfw.Press {
		param.yaw += speed
	}
	if win.GetKey(glfw.KeyRight) == glfw.Press {
		param.yaw -= speed
	}

	if win.GetKey(glfw.KeyQ) == glfw.Press {
		param.roll += speed
		if param.roll >= 360 {
			param.roll -= 360
		}
	}
	if win.GetKey(glfw.KeyE) == glfw.Press {
		param.roll -= speed
		if param.roll < 0 {
			param.roll += 360
		}
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
