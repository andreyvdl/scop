package render

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

func FBSCallback(w *glfw.Window, wid int, hei int) {
	setViewport(wid, hei)
}
