package main

import (
	"fmt"
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

	boid1 := Boid{0, 0, 0}
	fmt.Printf("%s\n", boid1.to_string())

	s.SetContent(0, 0, 'd', []rune{' ', 'd', ' '}, boxStyle)
	s.SetContent(0, 0, 'u', []rune{' '}, boxStyle)
	s.SetContent(0, 0, 'd', []rune{' '}, boxStyle)
	s.SetContent(0, 0, 'e', []rune{' '}, boxStyle)

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
