package main

import (
	"fmt"
)

type Boid struct {
	angle float32
	x float32
	y float32
}

func (boid Boid) to_string() string {
	message := fmt.Sprintf("Boid at position %f,%f with angle %f", boid.x, boid.y, boid.angle);
	return message
}