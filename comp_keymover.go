package main

import "github.com/hajimehoshi/ebiten"

type keyMover struct {
	el    *element
	speed float64
}

func (km *keyMover) update() error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		km.el.x -= km.speed
		km.el.rot = 0
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		km.el.x += km.speed
		km.el.rot = 180
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		km.el.y += km.speed
		km.el.rot = 90
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		km.el.y -= km.speed
		km.el.rot = 270
	}
	return nil
}

func (e *element) newkeyMover(speed float64) {
	km := &keyMover{e, speed}
	e.addComponent(km)

}
