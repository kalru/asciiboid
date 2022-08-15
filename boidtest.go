package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorReset)

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.EnablePaste()
	s.Clear()

	quit := func() {
		s.Fini()
		os.Exit(0)
	}

	x_max, y_max := s.Size()

	boids := []*Boid{}
	for i := 0; i < 50; i++ {
		boids = append(boids, &Boid{rand.Float64() * float64(x_max), rand.Float64() * float64(y_max),
			rand.Float64()*40 - 20, rand.Float64()*40 - 20, rand.Float64() * 20})
	}
	simulation := Simulation{boids}

	// refresh screen
	go func() {
		tickerRefresh := time.NewTicker(25 * time.Millisecond)
		for range tickerRefresh.C {
			s.Show()
		}
	}()

	// simulation update ticks
	go func() {
		tickRate := 25
		tickerUpdate := time.NewTicker(time.Duration(tickRate) * time.Millisecond)
		for range tickerUpdate.C {
			// update boids on screen
			simulation.simulate(float64(tickRate), s)

			// draw borders
			for _, i := range []int{1, x_max - 2} {
				for j := 1; j < y_max-1; j++ {
					s.SetContent(i, j, '|', nil, boxStyle)
				}
			}
			for _, j := range []int{0, y_max - 1} {
				for i := 0; i < x_max; i++ {
					s.SetContent(i, j, '-', nil, boxStyle)
				}
			}
		}
	}()

	for {
		// Poll event
		ev := s.PollEvent()
		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			x_max, y_max = s.Size()
			s.Clear()
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			}
		}
	}

}
