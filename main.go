package main

import (
	"fmt"
	"os"

	scop "scop/src"
	"scop/src/wavefront"

	glfw "github.com/go-gl/glfw/v3.3/glfw"
)

// func init() {
// 	runtime.LockOSThread()
// }

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, r)
		}
	}()

	if err := scop.ValidInput(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	wf := wavefront.NewWavefront()
	wf.LoadData(os.Args[1])

	startRendering(wf)
}

func startRendering(wf *wavefront.Wavefront) {
	_ = wf

	fmt.Println("{INFO|MAIN} starting window render")
	defer fmt.Println("{INFO|MAIN} window render stopped")

	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	win, err := glfw.CreateWindow(
		glfw.GetPrimaryMonitor().GetVideoMode().Width/2,
		glfw.GetPrimaryMonitor().GetVideoMode().Height/2,
		"scop",
		nil,
		nil,
	)
	if err != nil {
		panic(err)
	}

	win.MakeContextCurrent()

	for !win.ShouldClose() {
		win.SwapBuffers()
		glfw.PollEvents()
	}

}
