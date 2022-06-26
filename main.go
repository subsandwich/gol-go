package main

import (
	"time"
)

func main() {
	// Create new Game
	game := NewStandardGameOfLife([100][100]bool{{}, {false, false, true}, {true, false, true}, {false, true, true}})
	ui := NewGameOfLifeUI()
	ticker := time.NewTicker(200 * time.Millisecond)
	game.iterate()
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