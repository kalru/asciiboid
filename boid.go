package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Boid struct {
	x            float64
	y            float64
	dx           float64
	dy           float64
	visionRadius float64
}

func (boid Boid) to_string() string {
	message := fmt.Sprintf("Boid at position %f,%f", boid.x, boid.y)
	return message
}

func (boid *Boid) run(millis float64) {
	boid.x += boid.dx * float64(millis/float64(1000))
	boid.y += boid.dy * float64(millis/float64(1000))
}

func (boid Boid) render(s tcell.Screen) {
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorReset)
	s.SetContent(int(boid.x), int(boid.y), 'o', nil, boxStyle)
}

func (boid Boid) clear(s tcell.Screen) {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetContent(int(boid.x), int(boid.y), ' ', nil, defStyle)
}

func (boid *Boid) update(millis float64, s tcell.Screen) {
	// clear current pos
	boid.clear(s)
	// update values
	boid.run(millis)
	// redraw
	boid.render(s)
}
