package main

import (
	"fmt"
	"os"
	"runtime"

	scop "scop/src"
	"scop/src/render"
	"scop/src/wavefront"
)

func init() {
	runtime.LockOSThread()
}

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

	render.StartRendering(wf)
}
