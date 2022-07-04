package main

import (
	"testing"
	//	"fmt"
)

func Test_StandardGameOfLifeRules(t *testing.T) {
	data := []struct {
		name     string
		board    CellularAutomatonBoard
		x        int
		y        int
		expected bool
	}{
		{"wrap", [][]bool{{true, false, true}, {false, false, false}, {false, false, true}}, 0, 0, true},
		{"no neighbors", [][]bool{{false, false, false}, {false, false, false}, {false, false, false}}, 1, 1, false},
		{"1 neighbor", [][]bool{{true, false, false}, {false, true, false}, {false, false, false}}, 1, 1, false},
		{"2 neighbors", [][]bool{{true, true, false}, {false, true, false}, {false, false, false}}, 1, 1, true},
		{"2 neighbors differing rows", [][]bool{{false, false, false}, {true, true, false}, {false, false, true}}, 1, 1, true},
		{"3 neighbors", [][]bool{{true, true, false}, {false, true, true}, {false, false, false}}, 1, 1, true},
		{"4 neighbors", [][]bool{{true, true, false}, {false, true, true}, {true, false, false}}, 1, 1, false},
		{"revive", [][]bool{{true, false, false}, {false, false, true}, {true, false, false}}, 1, 1, true},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			if StandardGameOfLifeRules(d.board, d.x, d.y) != d.expected {
				t.Error("Invalid boolean")
			}
		})
	}
}

func Test_StandardGameOfLifeIterator(t *testing.T) {
	expectedIteratedBoard := [][]bool{{false, true, false, false, false}, {false, true, false, false, false}, {false, true, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}}
	first := StandardCellularAutomatonIterator([][]bool{{false, false, false, false, false}, {true, true, true, false, false}, {false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}}, StandardGameOfLifeRules)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if first[i][j] != expectedIteratedBoard[i][j] {
				t.Error("Invalid iteration")
				return
			}
		}
	}

}
