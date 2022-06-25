package main

import (
	"testing"
//	"fmt"
)

func Test_StandardGameOfLifeRules(t *testing.T) {
	wrapTest1 := [100][100]bool{{true}}
	wrapTest1[99][99] = true
	wrapTest1[99][0] = true
	data := []struct {
		name string
		board [100][100]bool
		x int
		y int
		expected bool
	}{
		{"no neighbors", [100][100]bool{}, 1, 1, false},
		{"1 neighbor", [100][100]bool{{true}, {false, true}}, 1, 1, false},
		{"2 neighbors", [100][100]bool{{true, true},{false, true}}, 1, 1, true},
		{"2 neighbors differing rows", [100][100]bool{{false},{true, true}, {false, false, true}}, 1, 1, true},
		{"3 neighbors", [100][100]bool{{true, true},{false, true, true}}, 1, 1, true},
		{"4 neighbors", [100][100]bool{{true, true},{false, true, true}, {true}}, 1, 1, false},
		{"revive", [100][100]bool{{true},{false, false, true}, {true}}, 1, 1, true},
		{"wrap live", wrapTest1, 0, 0, true},
		{"wrap 2", [100][100]bool{{true, true, true}}, 99, 1, true},
	}

	for _, d := range data {
		t.Run(d.name, func (t *testing.T){
			if StandardGameOfLifeRules(d.board, d.x, d.y) != d.expected {
				t.Error("Invalid boolean")
			}
		})
	}
}

func Test_StandardGameOfLifeIterator(t *testing.T) {
	expectedIteratedBoard := [100][100]bool{{false, true}, {false, true}}
	expectedIteratedBoard[99][1] = true
	first := StandardGameOfLifeIterator([100][100]bool{{true, true, true}})
	if first != expectedIteratedBoard {
		t.Error("Invalid iteration")
	}
}