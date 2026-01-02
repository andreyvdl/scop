package render

import (
	. "scop/src/render/matrices"
	. "scop/src/render/transformations"
	. "scop/src/render/vectors"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func KeyboardInput(win *glfw.Window, m *Mat4, degree *float64, scale *float32) {
	if win.GetKey(glfw.KeyEscape) == glfw.Press {
		win.SetShouldClose(true)
	}

	if win.GetKey(glfw.KeyZ) == glfw.Press {
		*scale += 0.001
		if *scale >= 1.0 {
			*scale = 1.0
		}
		*m = m.Mul(Scale(Vec3{*scale, *scale, *scale}))
	} else if win.GetKey(glfw.KeyX) == glfw.Press {
		*scale -= 0.001
		if *scale <= 0.5 {
			*scale = 0.5
		}
		*m = m.Mul(Scale(Vec3{*scale, *scale, *scale}))
	}

	if win.GetKey(glfw.KeyA) == glfw.Press {
		*degree -= 1.0
		if *degree <= 0 {
			*degree = 360.0
		}
		*m = m.Mul(Rotate(Deg2Rad(*degree), Vec3{0.0, 1.0, 0.0}))
	} else if win.GetKey(glfw.KeyD) == glfw.Press {
		*degree += 1.0
		if *degree >= 360.0 {
			*degree = 0.0
		}
		*m = m.Mul(Rotate(Deg2Rad(*degree), Vec3{0.0, 1.0, 0.0}))
	} else if win.GetKey(glfw.KeyW) == glfw.Press {
		*degree += 1.0
		if *degree >= 360.0 {
			*degree = 0.0
		}
		*m = m.Mul(Rotate(Deg2Rad(*degree), Vec3{0.0, 0.0, 1.0}))
	} else if win.GetKey(glfw.KeyS) == glfw.Press {
		*degree -= 1.0
		if *degree <= 0 {
			*degree = 360.0
		}
		*m = m.Mul(Rotate(Deg2Rad(*degree), Vec3{0.0, 0.0, 1.0}))
	} else if win.GetKey(glfw.KeyQ) == glfw.Press {
		*degree -= 1.0
		if *degree <= 0 {
			*degree = 360.0
		}
		*m = m.Mul(Rotate(Deg2Rad(*degree), Vec3{1.0, 0.0, 0.0}))
	} else if win.GetKey(glfw.KeyE) == glfw.Press {
		*degree += 1.0
		if *degree >= 360.0 {
			*degree = 0.0
		}
		*m = m.Mul(Rotate(Deg2Rad(*degree), Vec3{1.0, 0.0, 0.0}))
	}
}
