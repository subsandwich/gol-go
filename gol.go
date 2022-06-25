package main

// Iterator function for GOL
type GameOfLifeIterator func([100][100]bool) [100][100]bool


type GOL interface {
	iterate()
	newGame()
}

type GameOfLife struct {
	iterator GameOfLifeIterator
	iterations int
	board [100][100]bool
}

func (*g GameOfLife) iterate() {
	g.board = g.iterator(g.board)
	g.iterations += 1
}

func (*g GameOfLife) newGame() {
	g.board = [100][100]bool
	g.iterations = 0
	g.iterator = StandardGameOfLifeIterator
}


func StandardGameOfLifeIterator(b [100][100]bool) [100][100]bool {
	var newGOL [100][100]bool
	for i := 0; i < 100; i++{
		for j := 0; j < 100; j++{
			newGOL[i][j] = StandardGameOfLifeRules(newGOL, i, j)
		}
	}

}

func StandardGameOfLifeRules(b [100][100]bool, x int, y int) bool {
	neighborCount := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			posX := x+i
			posY := y+j
			if posX == x && posY == y {
				continue
			}
			if posX < 0 {
				posX = 100 + posX
			}
			if posY < 0 {
				posY = 100 + posY
			}
			if b[posX][posY]{
				neighborCount += 1
			}
		}
	}
	if !b[x][y]{
		return neighborCount == 3
	}
	return neighborCount == 2 || neighborCount == 3
}