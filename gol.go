package main

type CellularAutomatonBoard [][]bool

type CellularAutomatonRule func(CellularAutomatonBoard, int, int) bool

// Iterator function for GOL
type CellularAutomatonIterator func(CellularAutomatonBoard, CellularAutomatonRule) CellularAutomatonBoard

type ICellularAutomaton interface {
	iterate()
	reset()
}

type CellularAutomaton struct {
	rules      CellularAutomatonRule
	iterator   CellularAutomatonIterator
	iterations int
	board      CellularAutomatonBoard
}

func (g *CellularAutomaton) iterate() {
	g.board = g.iterator(g.board, g.rules)
	g.iterations++
}

func (g *CellularAutomaton) reset() {
	g.board = CellularAutomatonBoard{}
	g.iterations = 0
}

func NewStandardGameOfLife(b CellularAutomatonBoard) CellularAutomaton {
	return CellularAutomaton{
		StandardGameOfLifeRules,
		StandardCellularAutomatonIterator,
		0,
		b}
}

func StandardCellularAutomatonIterator(b CellularAutomatonBoard, r CellularAutomatonRule) CellularAutomatonBoard {
	board := make([][]bool, len(b))
	for i := 0; i < len(b); i++ {
		board[i] = make([]bool, len(b[i]))
		for j := 0; j < len(b[i]); j++ {
			board[i][j] = r(b, i, j)
		}
	}
	return board

}

func StandardGameOfLifeRules(b CellularAutomatonBoard, x int, y int) bool {
	neighborCount := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			posX := x + j
			posY := y + i
			if posX == x && posY == y {
				continue
			}

			if posX > len(b)-1 {
				posX = len(b) - posX
			} else if posX < 0 {
				posX = len(b) + posX
			}

			if posY > len(b[posX])-1 {
				posY = len(b[posX]) - posY
			} else if posY < 0 {
				posY = len(b[posX]) + posY
			}

			if b[posX][posY] {
				neighborCount++
			}
		}
	}
	if !b[x][y] {
		return neighborCount == 3
	}
	return neighborCount == 2 || neighborCount == 3
}
