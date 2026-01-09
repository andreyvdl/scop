package render

import "github.com/go-gl/glfw/v3.3/glfw"

func createWindow() *glfw.Window {
	video := glfw.GetPrimaryMonitor().GetVideoMode()
	width := video.Width / 2
	height := video.Height / 2

	win, err := glfw.CreateWindow(width, height, "scop", nil, nil)
	if err != nil {
		panic(err)
	}

	return win
}
