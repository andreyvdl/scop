package render

import "github.com/go-gl/glfw/v3.3/glfw"

func KeyboardInput(win *glfw.Window) {
	if win.GetKey(glfw.KeyEscape) == glfw.Press {
		win.SetShouldClose(true)
	}
}
