package main

import (
	"log"
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

	// i := 0
	// s.SetContent(i, 0, 'd', nil, boxStyle)

	// boid1 := Boid{0, 0, 0}
	// fmt.Printf("%s\n", boid1.to_string())

	boids := [1]Boid{}
	for i := 0; i < 1; i++ {
		boids[i] = Boid{2, 2, 10, 10, 10}
	}

	// refresh screen
	go func() {
		tickerRefresh := time.NewTicker(25 * time.Millisecond)
		for range tickerRefresh.C {
			s.Show()
		}
	}()

	// simulation update ticks
	go func() {
		tickerUpdate := time.NewTicker(100 * time.Millisecond)
		for range tickerUpdate.C {
			// update boids on screen
			for i := 0; i < 1; i++ {
				boids[i].update(100, s)
			}

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
