package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 800
	screenHeight = 800
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println(err)
		return
	}

	window, wErr := sdl.CreateWindow(
		"gaming",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)

	if wErr != nil {
		fmt.Println(wErr)
		return
	}
	defer window.Destroy()

	renderer, rErr := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if rErr != nil {
		fmt.Println(wErr)
		return
	}
	defer renderer.Destroy()

	sdl.LoadBMP

	for {
		renderer.SetDrawColor(255, 0, 255, 255)
		renderer.Clear()
		renderer.Present()
	}

}
