package main

import (
	"fmt"

	"github.com/T0PC4T/GoZombie/assets"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 800
)

func init() {

}

func update(screen *ebiten.Image) error {
	// Draws selected sprite image

	for _, e := range allElements {
		for _, c := range e.components {
			c.update()
		}
		for _, d := range e.drawers {
			d.draw(screen)
		}
	}
	// TPS counter
	fps := fmt.Sprintf("TPS: %f", ebiten.CurrentTPS())
	ebitenutil.DebugPrint(screen, fps)

	return nil
}

func main() {
	v := MyVector{x: 10, y: 0}
	fmt.Print(v.getRotation())
	player := newElement()
	player.newSpriteRenderer(assets.PlayerPng)
	player.newkeyMover(5.0)

	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "GoZombie"); err != nil {
		panic(err)
	}
}
