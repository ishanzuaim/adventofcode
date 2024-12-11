package main

import (
	"fmt"
	"strings"
)

type Position struct {
	x, y int
}

type GameSettings struct {
	grid [][]int
}

func main() {
	strs := readFile()

	gs := &GameSettings{}
	grid := make([][]int, len(strs))
	for y, row := range strs {
		grid[y] = convAll(strings.Split(row, ""))
	}
	gs.grid = grid

  zeroPos := gs.findZeros()
  total := 0
  for _, val := range zeroPos {
    check := gs.traverse(0, val)
    total += len(getTrueNinesP2(check))
  }
  // __AUTO_GENERATED_PRINT_VAR_START__
  fmt.Println(fmt.Sprintf("main total: %v", total)) // __AUTO_GENERATED_PRINT_VAR_END__
}

func getTrueNinesP2(pos []Position) []Position {
  nums := []Position{}
  for _, val := range pos {
    if val.x != -1  {
      nums = append(nums, val)
    }
  }
  return nums
}


func getTrueNines(pos []Position) []Position {
  nums := []Position{}
  for _, val := range pos {
    if val.x != -1 && !contains(val, nums) {
      nums = append(nums, val)
    }
  }
  return nums
}


func contains(val Position, pos []Position) bool {
  for _, ival := range pos {
    if ival.x == val.x && ival.y == val.y {
      return true
    }
  }
  return false
}

func (gs *GameSettings) findZeros() []Position {
  total := []Position{}
  for y, row := range gs.grid {
    for x, val := range row {
      if val == 0 {
        total = append(total, Position{x, y})
      }
    }
  }
  return total;
}

func (gs *GameSettings) traverse(idx int, pos Position) []Position {
	if gs.grid[pos.y][pos.x] != idx {
		return []Position{{-1, -1}}
	}

	if gs.grid[pos.y][pos.x] == 9 {
		return []Position{{pos.x, pos.y}}
	}

  total := []Position{}

	if pos.x < len(gs.grid[0]) - 1 {
		total = append(total, gs.traverse(idx+1, Position{pos.x + 1, pos.y})...)
	}
	if pos.x > 0 {
		total = append(total, gs.traverse(idx+1, Position{pos.x - 1, pos.y})...)
	}
	if pos.y < len(gs.grid) - 1 {
		total = append(total, gs.traverse(idx+1, Position{pos.x, pos.y + 1})...)
	}
	if pos.y > 0 {
		total = append(total, gs.traverse(idx+1, Position{pos.x, pos.y - 1})...)
	}
	return total
}
