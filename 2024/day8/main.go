package main

import (
	"fmt"
	"strings"
  "github.com/ishanzuaim/adventofcode/util"
)

type Position struct {
	x int
	y int
}

func (posA Position) isEqual(posB Position) bool {
	return posA.x == posB.x && posA.y == posB.y
}

func (posA Position) isOutsideBounds(boundaries Position) bool {
	return posA.x < 0 || posA.y < 0 || posA.y >= boundaries.y || posA.x >= boundaries.x
}

func main() {
	strs := readFile()
	gs := &GameSettings{}
	gs.init(strs)
	var all []Position = make([]Position, 0)
	for key := range gs.freqMap {
		occs := gs.findCloseAntinodesP2(key)
		all = merge(all, occs)
	}
	// __AUTO_GENERATED_PRINT_VAR_START__
	fmt.Println(fmt.Sprintf("main all: %v", len(all))) // __AUTO_GENERATED_PRINT_VAR_END__

	for y := 0; y < gs.boundaries.y; y++ {
		for x := 0; x < gs.boundaries.x; x++ {
			if isExist(all, Position{x, y}) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func merge(posA, posB []Position) []Position {
	newPos := posA
	for _, val := range posB {
		if !isExist(newPos, val) {
			newPos = append(newPos, val)
		}
	}
	return newPos
}

type GameSettings struct {
	freqMap    map[string][]Position
	boundaries Position
}

func (gs *GameSettings) init(strs []string) {
	gs.freqMap = make(map[string][]Position)
	gs.boundaries.x = len(strs)
	for y, row := range strs {
		cols := strings.Split(row, "")
		gs.boundaries.y = len(cols)
		for x, elem := range cols {
			if elem != "." {
				_, ok := gs.freqMap[elem]
				if ok {
					gs.freqMap[elem] = append(gs.freqMap[elem], Position{x, y})
				} else {
					gs.freqMap[elem] = []Position{{x, y}}
				}
			}
		}
	}
}

func (gs *GameSettings) findCloseAntinodesP2(chr string) []Position {
	var close []Position
	occs := gs.freqMap[chr]
	if occs == nil {
		panic("Should have a char")
	}
	for idxA, posA := range occs {
		for idxB, posB := range occs {
			if idxA == idxB {
				continue
			}
			tempA := posA
			tempB := posB
			for true {
				pos := getCloseAnitnode(tempA, tempB)
				if !isExist(close, tempA) {
					close = append(close, tempA)
				}
				if !pos.isOutsideBounds(gs.boundaries) && !isExist(close, pos) {
					close = append(close, pos)
				} else {
					break
				}
				tempA, tempB = pos, tempA
			}
		}
	}
	return close
}

func (gs *GameSettings) findCloseAntinodes(chr string) []Position {
	var close []Position
	occs := gs.freqMap[chr]
	if occs == nil {
		panic("Should have a char")
	}
	for idxA, posA := range occs {
		for idxB, posB := range occs {
			if idxA == idxB {
				continue
			}
			pos := getCloseAnitnode(posA, posB)
			if !pos.isOutsideBounds(gs.boundaries) && !isExist(close, pos) {
				close = append(close, pos)
			}
		}
	}
	return close
}

func isExist(positions []Position, pos Position) bool {
	for _, val := range positions {
		if val.isEqual(pos) {
			return true
		}
	}
	return false
}

func getCloseAnitnode(posA, posB Position) Position {
	diffX, diffY := posB.x-posA.x, posB.y-posA.y
	return Position{posA.x - diffX, posA.y - diffY}
}
