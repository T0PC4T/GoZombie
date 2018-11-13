package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type component interface {
	onUpdate() error
	onDraw(renderer *sdl.Renderer) error
}

type element struct {
	x          float64
	y          float64
	rot        float64
	active     bool
	components [](*component)
}

func (elem *element) addComponent(c component) {
	elem.components = append(elem.components, &c)
	// if err != nil {
	// 	fmt.Println("THERE WAS AN ERROR IN APPEND", err)
	// }
}

var allElements [](*element)
