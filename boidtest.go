package main

import (
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)

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

	i := 1

	s.SetContent(i, 0, 'd', nil, boxStyle)

	// boid1 := Boid{0, 0, 0}
	// fmt.Printf("%s\n", boid1.to_string())

	boids := [10]Boid{}
	for i := 0; i < 10; i++ {
		boids[i] = Boid{float32(i), float32(i), 0, 10}
	}

	// refresh screen
	go func() {
		tickerRefresh := time.NewTicker(50 * time.Millisecond)
		for _ = range tickerRefresh.C {
			s.Show()
		}
	}()

	// simulation update ticks
	go func() {
		tickerUpdate := time.NewTicker(100 * time.Millisecond)
		for _ = range tickerUpdate.C {
			// TODO only call update method of each boid here passing a time param in millis
			s.SetContent(i, 0, ' ', nil, defStyle)
			i += 1
			s.SetContent(i, 0, 'd', nil, boxStyle)
		}
	}()

	for {
		// Poll event
		ev := s.PollEvent()
		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			}
		}
	}

}
