package main

import (
	"log"
	"os"

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

	// boid1 := Boid{0, 0, 0}
	// fmt.Printf("%s\n", boid1.to_string())

	boids := [10]Boid{}
	for i := 0; i < 10; i++ {
		boids[i] = Boid{float32(i), float32(i), 0, 10}
	}

	s.SetContent(0, 0, 'd', nil, boxStyle)
	
	for {
		// Update screen
		s.Show()
		// Poll event
		ev := s.PollEvent()
		// Process event
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			}
		}
	}

}
