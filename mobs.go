package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	zombieSpeed       float64 = 0.02
	zombieImgWidth    int32   = 313
	zombieImgHeight   int32   = 207
	zombieScaleWidth  int32   = 150
	zombieScaleHeight int32   = 100
)

type zombie struct {
	texture *sdl.Texture
	x, y    float64
	speed   float64
}

func NewZombie(renderer *sdl.Renderer) (z zombie, err error) {
	z.speed = playerSpeed
	z.x = 200
	z.y = 200

	zombieImg, err := sdl.LoadBMP("assets/zombie.bmp")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer zombieImg.Free()

	z.texture, err = renderer.CreateTextureFromSurface(zombieImg)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func (z *zombie) draw(renderer *sdl.Renderer) {
	renderer.Copy(z.texture,
		&sdl.Rect{X: 0, Y: 0, W: zombieImgWidth, H: zombieImgHeight},
		&sdl.Rect{X: int32(z.x), Y: int32(z.y), W: zombieScaleWidth, H: zombieScaleHeight})
}

func (z *zombie) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_UP] == 1 {
		z.x -= z.speed
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		z.x += z.speed
	}
}
