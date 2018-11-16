package main

import (
	"bytes"
	"image"

	"github.com/T0PC4T/GoZombie/assets"
	"github.com/hajimehoshi/ebiten"
)

type spriteRenderer struct {
	el  *element
	img *ebiten.Image
}

func (sr *spriteRenderer) draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(float64(sr.el.x), float64(sr.el.y))
	op.GeoM.Rotate(sr.el.rot)

	screen.DrawImage(sr.img, op)
	return nil
}

func (e *element) newSpriteRenderer(bslice []byte) {
	// Preload images
	img, _, err := image.Decode(bytes.NewReader(assets.PlayerPng))
	if err != nil {
		panic(err)
	}

	fullImg, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	sr := &spriteRenderer{img: fullImg, el: e}
	e.addDrawer(sr)
}
