package main

import "math"

type MyVector struct {
	x float64
	y float64
}

func (v *MyVector) getRotation() float64 {
	return math.Atan2(v.x, v.y)
}
