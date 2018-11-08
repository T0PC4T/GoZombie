package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	texture *sdl.Texture
}

func newPlayer(renderer *sdl.Renderer) (p player, err error) {

	playerImg, err := sdl.LoadBMP("assets/player.bmp")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer playerImg.Free()

	p.texture, err = renderer.CreateTextureFromSurface(playerImg)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func (p *player) draw(renderer *sdl.Renderer) {
	renderer.Copy(p.texture,
		&sdl.Rect{X: 0, Y: 0, W: 300, H: 400},
		&sdl.Rect{X: 0, Y: 0, W: 100, H: 100})
}
