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

	window, err := sdl.CreateWindow(
		"gaming",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()
	// Create Player
	err = newPlayer(renderer)
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer myPlayer.texture.Destroy()
	// Create Zombie

	// myMob, err := NewZombie(renderer)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer myMob.texture.Destroy()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 0, 255, 255)
		renderer.Clear()

		for _, e := range allElements {
			for _, c := range e.components {
				(*c).onUpdate()
				(*c).onDraw(renderer)
			}
		}

		renderer.Present()
	}

}
