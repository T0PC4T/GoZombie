package main

import (
	"github.com/hajimehoshi/ebiten"
)

type component interface {
	update() error
}

type drawer interface {
	draw(*ebiten.Image) error
}

type element struct {
	x          float64
	y          float64
	rot        float64
	active     bool
	components []component
	drawers    []drawer
}

// adds a component to the element components
func (elem *element) addComponent(c component) {
	elem.components = append(elem.components, c)
	// if err != nil {
	// 	fmt.Println("THERE WAS AN ERROR IN APPEND", err)
	// }
}

// adds a component to the element components
func (elem *element) addDrawer(d drawer) {
	elem.drawers = append(elem.drawers, d)
	// if err != nil {
	// 	fmt.Println("THERE WAS AN ERROR IN APPEND", err)
	// }
}

func newElement() *element {
	e := &element{}
	allElements = append(allElements, e)
	return e
}

var allElements [](*element)
