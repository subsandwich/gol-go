package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

func main() {
	// Create new Game
	board := make([][]bool, 20)
	for i := range board {
		board[i] = make([]bool, 20)
	}
	board[0][2] = true
	board[1][0] = true
	board[1][2] = true
	board[2][1] = true
	board[2][2] = true
	game := NewStandardGameOfLife(board)
	ui := NewCellularAutomatonUI()
	defer ui.screen.Fini()
	ticker := time.NewTicker(200 * time.Millisecond)
	ui.update(game)
	go func() {
		for {

			select {
			case <-ticker.C:
				if !ui.paused {
					game.iterate()
					ui.update(game)
				}

			}
		}
	}()

	for {
		event := ui.screen.PollEvent()

		switch event := event.(type) {
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEnter {
				ui.togglePause()
			}
			if event.Key() == tcell.KeyEsc {
				return
			}

		case *tcell.EventMouse:
			x, y := event.Position()

			switch event.Buttons() {
			case tcell.Button1:
				ui.flipBit(x, y, game)
			}
		}

	}

}
