package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	el  *element
	tex *sdl.Texture
}

func (sr *spriteRenderer) onUpdate() error {
	fmt.Println("updating")
	return nil
}

func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error {
	fmt.Println("Drawing")
	renderer.CopyEx(sr.tex,
		&sdl.Rect{X: 0, Y: 0, W: playerImgWidth, H: playerImgHeight},
		&sdl.Rect{X: int32(sr.el.x), Y: int32(sr.el.y), W: playerScaleWidth, H: playerScaleHeight},
		sr.el.rot,
		&sdl.Point{X: playerScaleWidth / 2, Y: playerScaleHeight / 2},
		sdl.FLIP_NONE)
	return nil
}

func newSpriteRenderer(renderer *sdl.Renderer, filename string) *spriteRenderer {

	img, err := sdl.LoadBMP(filename)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer img.Free()

	texture, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return &spriteRenderer{tex: texture}

}
