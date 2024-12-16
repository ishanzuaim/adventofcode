package main

import (
	"fmt"
	"strings"
)

type GameSettings struct {
	state map[Position][]int
	grid  [][]string
}

type Position struct {
	x int
	y int
}

func main() {
	strs := readFile()
	gs := &GameSettings{
		make(map[Position][]int),
		make([][]string, len(strs)),
	}
	gs.init(strs)
  gs.calcState()
  price := gs.calcPrice()
  // __AUTO_GENERATED_PRINT_VAR_START__
  fmt.Println(fmt.Sprintf("main price: %v", price)) // __AUTO_GENERATED_PRINT_VAR_END__
}

func (gs *GameSettings) init(strs []string) {
	for y, row := range strs {
		line := strings.Split(row, "")
		gs.grid[y] = line
	}
}

func (gs *GameSettings) calcState() {
  idx := "-1"
	for y, row := range gs.grid {
		for x, val := range row {
			if !strings.Contains(val, "-") {
				pos := Position{x, y}
				a, b := gs.capture(val,idx, pos)
				gs.state[pos] = []int{a, b}
        // __AUTO_GENERATED_PRINT_VAR_START__
        fmt.Println(fmt.Sprintf("calcState gs: %v", gs.grid)) // __AUTO_GENERATED_PRINT_VAR_END__
        idx = convToStr(convToInt(idx) - 1)
			}
		}
	}
}

func (gs *GameSettings) capture(val, idx string, pos Position) (int, int) {
	if pos.x < 0 || pos.x >= len(gs.grid[0]) {
		return 0, 1
	}
	if pos.y < 0 || pos.y >= len(gs.grid) {
		return 0, 1
	}

	checkVal := gs.grid[pos.y][pos.x]
	if checkVal != val {
    if checkVal == idx {
      return 0, 0
    }
		return 0, 1
	}
	gs.grid[pos.y][pos.x] = idx
	totalA := 1
	totalB := 0
	ta, tp := gs.capture(val, idx, Position{pos.x + 1, pos.y})
	totalA += ta
	totalB += tp
	ta, tp = gs.capture(val, idx,Position{pos.x - 1, pos.y})
	totalA += ta
	totalB += tp
	ta, tp = gs.capture(val, idx, Position{pos.x, pos.y + 1})
	totalA += ta
	totalB += tp
	ta, tp = gs.capture(val, idx, Position{pos.x, pos.y - 1})
	totalA += ta
	totalB += tp
	return totalA, totalB
}

func (gs *GameSettings) calcPrice() int {
	total := 0
	for _, val := range gs.state {
		total += val[0] * val[1]
	}
	return total
}

func (gs *GameSettings) checkDir(pos Position) int {
	val := gs.grid[pos.y][pos.x]
	dirs := []int{-1, 1}

	perim := 0
	for _, dir := range dirs {
		checkPos := Position{pos.x + dir, pos.y}
		if checkPos.x < 0 {
			perim++
		} else if checkPos.x >= len(gs.grid[0]) {
			perim++
		} else {
			checkVal := gs.grid[checkPos.y][checkPos.x]
			if checkVal != val {
				perim++
			}
		}
	}

	for _, dir := range dirs {
		checkPos := Position{pos.x, pos.y + dir}
		if checkPos.y < 0 {
			perim++
		} else if checkPos.y >= len(gs.grid) {
			perim++
		} else {
			checkVal := gs.grid[checkPos.y][checkPos.x]
			if checkVal != val {
				perim++
			}
		}
	}
	return perim
}
