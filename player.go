package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed       float64 = 0.07
	playerImgWidth    int32   = 313
	playerImgHeight   int32   = 207
	playerScaleWidth  int32   = 150
	playerScaleHeight int32   = 100
)

func newPlayer(renderer *sdl.Renderer) error {
	p := new(element)
	// p.speed = playerSpeed
	p.x = 300
	p.y = 300
	sr := newSpriteRenderer(renderer, "assets/player.bmp")
	p.addComponent(sr)
	allElements = append(allElements, p)

	return nil
}

// func (p *player) update() {
// 	keys := sdl.GetKeyboardState()

// 	if keys[sdl.SCANCODE_RIGHT] == 1 {
// 		p.x += p.speed
// 		p.rot = 0
// 	} else if keys[sdl.SCANCODE_LEFT] == 1 {
// 		p.x -= p.speed
// 		p.rot = 180
// 	} else if keys[sdl.SCANCODE_DOWN] == 1 {
// 		p.y += p.speed
// 		p.rot = 90
// 	} else if keys[sdl.SCANCODE_UP] == 1 {
// 		p.y -= p.speed
// 		p.rot = 270
// 	}

// }
