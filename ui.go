package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

type ICellularAutomatonUI interface {
	update(CellularAutomaton)
	togglePause()
}

type CellularAutomatonUI struct {
	name   string
	screen tcell.Screen
	style  tcell.Style
	length int
	width  int
	x      int
	y      int
	offset int
	paused bool
}

func (ui CellularAutomatonUI) update(g CellularAutomaton) {
	tbr := tcell.StyleDefault.Background(tcell.ColorWhite)
	for x := 0; x < ui.width; x++ {
		for y := 0; y < ui.length; y++ {
			if g.board[x][y] {
				ui.screen.SetContent(x*2+ui.x+ui.offset*2, y+ui.y+ui.offset, ' ', nil, ui.style)
				ui.screen.SetContent(x*2+1+ui.x+ui.offset*2, y+ui.y+ui.offset, ' ', nil, ui.style)
			} else {
				ui.screen.SetContent(x*2+ui.x+ui.offset*2, y+ui.y+ui.offset, ' ', nil, tbr)
				ui.screen.SetContent(x*2+1+ui.x+ui.offset*2, y+ui.y+ui.offset, ' ', nil, tbr)
			}

		}
	}
	ui.screen.Show()
}

func (ui *CellularAutomatonUI) togglePause() {
	ui.paused = !ui.paused
	ui.drawInfo()
}

func NewCellularAutomatonUI() CellularAutomatonUI {
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
	s.EnableMouse()

	newCA := CellularAutomatonUI{"Conway's Game of Life", s, boxStyle, 20, 20, 1, 1, 1, false}
	newCA.drawSquare()
	newCA.drawInfo()
	s.Show()
	return newCA
}

func (ui CellularAutomatonUI) drawSquare() {
	// Double width
	sWidth := ui.width * 2

	perimStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

	// Make corners
	ui.screen.SetContent(ui.x, ui.y, tcell.RuneULCorner, nil, perimStyle)
	ui.screen.SetContent(ui.x+sWidth+ui.offset*4, ui.y, tcell.RuneURCorner, nil, perimStyle)
	ui.screen.SetContent(ui.x, ui.y+ui.length+ui.offset*2, tcell.RuneLLCorner, nil, perimStyle)
	ui.screen.SetContent(ui.x+sWidth+ui.offset*4, ui.y+ui.length+ui.offset*2, tcell.RuneLRCorner, nil, perimStyle)

	// Set Horizontal
	for x := ui.x + 1; x < ui.x+sWidth+ui.offset*4; x++ {
		ui.screen.SetContent(x, ui.y, tcell.RuneHLine, nil, perimStyle)
		ui.screen.SetContent(x, ui.y+ui.length+ui.offset*2, tcell.RuneHLine, nil, perimStyle)
	}

	// Set Vertical
	for y := ui.y + 1; y < ui.y+ui.length+ui.offset*2; y++ {
		ui.screen.SetContent(ui.x, y, tcell.RuneVLine, nil, perimStyle)
		ui.screen.SetContent(ui.x+sWidth+ui.offset*4, y, tcell.RuneVLine, nil, perimStyle)
	}

	// Set Background
	for x := ui.x + 1 + ui.offset*2; x < ui.x+sWidth+ui.offset*2; x++ {
		for y := ui.y + 1 + ui.offset*2; y < ui.y+ui.length+ui.offset*2; y++ {
			ui.screen.SetContent(x, y, ' ', nil, ui.style)
		}
	}

}

func (ui CellularAutomatonUI) drawInfo() {
	infoStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	infoY := ui.y + ui.offset*2 + ui.length + 1

	for i, rVal := range "Current Cellular Automaton: " + ui.name {
		ui.screen.SetContent(ui.x+i, infoY, rVal, nil, infoStyle)
	}

	pauseString := "running"
	if ui.paused {
		pauseString = "paused"
	}
	for i, rVal := range "Game is currently " + pauseString {
		ui.screen.SetContent(ui.x+i, infoY+1, rVal, nil, infoStyle)
	}
	ui.screen.Show()
}

func (ui CellularAutomatonUI) flipBit(x, y int, g CellularAutomaton) {
	tbr := tcell.StyleDefault.Background(tcell.ColorWhite)

	addInt := 0
	if (ui.x+ui.offset)%2 == 0 {
		addInt = 1
	}
	xIndex := (x+addInt)/2 - ui.x - ui.offset
	yIndex := y - ui.y - ui.offset
	if xIndex >= ui.width || yIndex >= ui.length || xIndex < 0 || yIndex < 0 {
		return
	}
	g.board[xIndex][yIndex] = !g.board[xIndex][yIndex]

	if x%2 == 0 {
		x--
	}
	if g.board[xIndex][yIndex] {
		ui.screen.SetContent(x, y, ' ', nil, ui.style)
		ui.screen.SetContent(x+1, y, ' ', nil, ui.style)
	} else {
		ui.screen.SetContent(x, y, ' ', nil, tbr)
		ui.screen.SetContent(x+1, y, ' ', nil, tbr)
	}

	ui.screen.Show()
}
