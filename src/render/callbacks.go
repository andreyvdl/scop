package render

import (
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func FBSCallback(w *glfw.Window, wid int, hei int) {
	gl.Viewport(0, 0, int32(wid), int32(hei))
}

func KCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}
}
