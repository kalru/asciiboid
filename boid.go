package main

import (
	"math"
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

type Boid struct {
	x            float64
	y            float64
	dx           float64
	dy           float64
	visionRadius float64
}

type Simulation struct {
	boids []*Boid
}

func (boid *Boid) run(millis float64, s tcell.Screen) {
	// limit speed
	boid.limitSpeed()

	// keep in bounds
	x_max, y_max := s.Size()
	var margin float64 = 15
	var turnFactor float64 = rand.Float64()
	if boid.x < margin {
		boid.dx += turnFactor
	} else if boid.x > float64(x_max)-margin {
		boid.dx -= turnFactor
	}
	if boid.y < margin {
		boid.dy += turnFactor
	} else if boid.y > float64(y_max)-margin {
		boid.dy -= turnFactor
	}

	// update pos based on velocity
	boid.x += boid.dx * float64(millis/float64(1000))
	boid.y += boid.dy * float64(millis/float64(1000))
}

func (boid Boid) render(s tcell.Screen) {
	speed := math.Sqrt(boid.dx*boid.dx + boid.dy*boid.dy)
	boidColor := tcell.PaletteColor(int(1 + int(speed)/5))
	boidStyle := tcell.StyleDefault.Foreground(boidColor).Background(tcell.ColorReset)
	angle := math.Atan2(boid.dy, boid.dx)
	var char rune
	if angle <= math.Pi/8 && angle > -math.Pi/8 {
		char = '→'
	} else if (angle <= math.Pi && angle > math.Pi*7/8) || angle >= -math.Pi && angle <= -math.Pi*7/8 {
		char = '←'
	} else if angle > math.Pi/8 && angle <= math.Pi*3/8 {
		char = '↘'
	} else if angle > -math.Pi*7/8 && angle <= -math.Pi*5/8 {
		char = '↖'
	} else if angle > math.Pi*3/8 && angle <= math.Pi*5/8 {
		char = '↓'
	} else if angle > -math.Pi*5/8 && angle <= -math.Pi*3/8 {
		char = '↑'
	} else if angle > math.Pi*5/8 && angle <= math.Pi*7/8 {
		char = '↙'
	} else if angle > -math.Pi*3/8 && angle <= -math.Pi/8 {
		char = '↗'
	}
	s.SetContent(int(math.Round(boid.x)), int(math.Round(boid.y)), char, nil, boidStyle)
}

func (boid Boid) clear(s tcell.Screen) {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetContent(int(math.Round(boid.x)), int(math.Round(boid.y)), ' ', nil, defStyle)
}

func (boid *Boid) update(millis float64, s tcell.Screen) {
	boid.clear(s)
	boid.run(millis, s)
	boid.render(s)
}

func dist(boid1 Boid, boid2 Boid) float64 {
	return math.Sqrt((boid1.x-boid2.x)*(boid1.x-boid2.x) + (boid1.y-boid2.y)*(boid1.y-boid2.y))
}

func (boid *Boid) flyTowardsCenter(boids []*Boid) {
	// find the center of mass of the other boids and adjust velocity slightly to
	// point towards the center of mass.
	centeringFactor := 0.1 // adjust velocity by this %
	centerX, centerY, numNeighbors := 0.0, 0.0, 0.0
	for _, b := range boids {
		if b != boid && dist(*boid, *b) < boid.visionRadius {
			centerX += b.x
			centerY += b.y
			numNeighbors++
		}
	}
	if numNeighbors > 0 {
		centerX = centerX / numNeighbors
		centerY = centerY / numNeighbors
		boid.dx += (centerX - boid.x) * centeringFactor
		boid.dy += (centerY - boid.y) * centeringFactor
	}
}

func (boid *Boid) avoidOthers(boids []*Boid) {
	minDist := 5.0     // The distance to stay away from other boids
	avoidFactor := 0.1 // Adjust velocity by this %
	moveX, moveY := 0.0, 0.0
	for _, b := range boids {
		if b != boid && dist(*boid, *b) < minDist {
			moveX += boid.x - b.x
			moveY += boid.y - b.y
		}
	}
	boid.dx += moveX * avoidFactor
	boid.dy += moveY * avoidFactor
}

func (boid *Boid) matchVelocity(boids []*Boid) {
	// Find the average velocity (speed and direction) of the other boids and
	// adjust velocity slightly to match.
	matchingFactor := 0.005 // Adjust by this % of average velocity
	avgDX, avgDY, numNeighbors := 0.0, 0.0, 0.0
	for _, b := range boids {
		if b != boid && dist(*boid, *b) < boid.visionRadius {
			avgDX += b.dx
			avgDY += b.dy
			numNeighbors++
		}
	}
	if numNeighbors > 0 {
		avgDX = avgDX / numNeighbors
		avgDY = avgDY / numNeighbors
		boid.dx += (avgDX - boid.dx) * matchingFactor
		boid.dy += (avgDY - boid.dy) * matchingFactor
	}
}

func (boid *Boid) limitSpeed() {
	// Speed will naturally vary in flocking behavior, but real animals can't go
	// arbitrarily fast.
	speedLimit := 30.0
	speed := math.Sqrt(boid.dx*boid.dx + boid.dy*boid.dy)
	if speed > speedLimit {
		// proportially decrese speed
		boid.dx = (boid.dx / speed) * speedLimit
		boid.dy = (boid.dy / speed) * speedLimit
	}
}

func (simulation *Simulation) averageSpeed() float64 {
	speed := 0.0
	for _, boid := range simulation.boids {
		speed += math.Sqrt(boid.dx*boid.dx + boid.dy*boid.dy)
	}
	return speed / float64(len(simulation.boids))
}

func (simulation *Simulation) simulate(millis float64, s tcell.Screen) {
	for _, boid := range simulation.boids {
		// This is very inefficient, but works for now
		boid.flyTowardsCenter(simulation.boids)
		boid.avoidOthers(simulation.boids)
		boid.matchVelocity(simulation.boids)
		boid.update(millis, s)
	}
}
