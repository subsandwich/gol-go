package main

import (
	"time"
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
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				game.iterate()
				ui.update(game)
			}
		}
	}()

	time.Sleep(10000 * time.Millisecond)
	ticker.Stop()
	done <- true

}
