package main

import (
	"github.com/gdamore/tcell/v2"
	"log"
)

type IGameOfLifeUI interface {
	update(GameOfLife)
}

type GameOfLifeUI struct {
	screen tcell.Screen
	style tcell.Style
}

func (ui GameOfLifeUI) update(g GameOfLife) {
	tbr := tcell.StyleDefault.Background(tcell.ColorWhite)
	ui.screen.Clear()
	for row := 0; row < 100; row++ {
		for col := 0; col < 100; col++ {
			if g.board[row][col] {
				ui.screen.SetContent(col*2, row, ' ', nil, ui.style)
				ui.screen.SetContent(col*2+1, row, ' ', nil, ui.style)
			} else {
				ui.screen.SetContent(col*2, row, ' ', nil, tbr)
				ui.screen.SetContent(col*2+1, row, ' ', nil, tbr)
			}

		}
	}
	ui.screen.Show()
}

func NewGameOfLifeUI() GameOfLifeUI {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.Clear()

	boxStyle := tcell.StyleDefault.Background(tcell.ColorBlack)
	s.SetStyle(boxStyle)
	// Fill background
	for row := 0; row < 100; row++ {
		for col := 0; col < 100; col++ {
			s.SetContent(col, row, ' ', nil, boxStyle)
		}
	}
	s.Show()
	return GameOfLifeUI{s, boxStyle}
}
