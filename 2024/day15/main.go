package main

import (
	"fmt"
	"strings"
	"time"
	// "time"
)

type GameSettings struct {
	grid [][]string
	dirs []string
	pos  Position
}

func main() {
	strs := readFile()
	gs := &GameSettings{}
	gs.init(strs)
	gs.navigateP2()
	val := gs.calcValueP2()
  // __AUTO_GENERATED_PRINT_VAR_START__
  fmt.Println(fmt.Sprintf("main val: %v", val)) // __AUTO_GENERATED_PRINT_VAR_END__
}

func (gs *GameSettings) print() {
  fmt.Printf("\033[0;0H")
  // fmt.Println(gs.grid)
	for _, row := range gs.grid {
		for _, val := range row {
			fmt.Print(val)
		}
		fmt.Print("\n")
	}
}

func (gs *GameSettings) init(strs []string) {
	flag := 0
	gs.grid = make([][]string, 0)
	gs.dirs = make([]string, 0)
	for idx, row := range strs {
		if row == "" {
			flag = 1
			continue
		}
		if flag == 0 {
			gs.grid = append(gs.grid, []string{})
			gs.grid[idx] = strings.Split(row, "")
		} else {
			gs.dirs = append(gs.dirs, strings.Split(row, "")...)
		}
	}
	gs.findRobot()
}

type Position struct {
	x, y int
}

func (gs *GameSettings) findRobot() {
	for y, row := range gs.grid {
		for x, val := range row {
			if val == "@" {
				gs.pos = Position{x, y}
				return
			}
		}
	}
	panic("Should have found robot")
}

func (gs *GameSettings) isBox(pos Position) bool {
	obj := gs.grid[pos.y][pos.x]
	return obj == "]" || obj == "["
}

func (gs *GameSettings) navigateP2() {
	m := map[string][2]int{
		"v": {0, 1},
		"<": {-1, 0},
		"^": {0, -1},
		">": {1, 0},
	}
	for _, val := range gs.dirs {
		dir, _ := m[val]
		nextPos := getNewPos(gs.pos, dir)
		time.Sleep(200 * time.Millisecond)
		if gs.moveBoxP2(nextPos, dir) {
			if dir[1] != 0 {
				gs.move(nextPos, dir)
			}
			gs.grid[gs.pos.y][gs.pos.x] = "."
			gs.pos = getNewPos(gs.pos, dir)
			gs.grid[gs.pos.y][gs.pos.x] = "@"
		}
    // __AUTO_GENERATED_PRINT_VAR_START__
    fmt.Println(fmt.Sprintf("navigateP2 val: %v", val)) // __AUTO_GENERATED_PRINT_VAR_END__
		gs.print()
	}
}

func getNewPos(pos Position, dir [2]int) Position {
	return Position{pos.x + dir[0], pos.y + dir[1]}
}

func (gs *GameSettings) moveBoxP2(pos Position, dir [2]int) bool {
	curr_obj := gs.grid[pos.y][pos.x]
	if curr_obj == "." {
		return true
	} else if curr_obj == "#" {
		return false
	}

	new_pos := getNewPos(pos, dir)
	if dir[0] != 0 {
		canMove := gs.moveBoxP2(new_pos, dir)
		if canMove {
			gs.grid[new_pos.y][new_pos.x] = curr_obj
			gs.grid[new_pos.y][pos.x] = "."
		}
		return canMove
	}

	if curr_obj == "]" {
		canMove := gs.moveBoxP2(new_pos, dir) &&
			gs.moveBoxP2(Position{new_pos.x - 1, new_pos.y}, dir)

		return canMove
	} else {
		canMove := gs.moveBoxP2(new_pos, dir) && gs.moveBoxP2(Position{new_pos.x + 1, new_pos.y}, dir)
		return canMove
	}
}

func (gs *GameSettings) calcValueP2() int {
	total := 0
	for y, row := range gs.grid {
		for x, val := range row {
			if val == "[" {
				total += getCoord(Position{x, y})
			}
		}
	}
	return total
}

func (gs *GameSettings) move(pos Position, dir [2]int) {
	curr := gs.grid[pos.y][pos.x]
	if curr == "." {
		return
	}

	new_pos := getNewPos(pos, dir)

	if curr == "]" {
		gs.move(Position{new_pos.x - 1, new_pos.y}, dir)
		gs.move(Position{new_pos.x, new_pos.y}, dir)
		gs.grid[new_pos.y][new_pos.x-1] = "["
		gs.grid[new_pos.y][new_pos.x] = "]"
		gs.grid[pos.y][pos.x] = "."
		gs.grid[pos.y][pos.x-1] = "."
	}
	if curr == "[" {
		gs.move(Position{new_pos.x + 1, new_pos.y}, dir)
		gs.move(Position{new_pos.x, new_pos.y}, dir)
		gs.grid[new_pos.y][new_pos.x] = "["
		gs.grid[new_pos.y][new_pos.x+1] = "]"
		gs.grid[pos.y][pos.x] = "."
		gs.grid[pos.y][pos.x+1] = "."
	}
	// if gs.grid[new_pos.y][new_pos.x] == "." && gs.g {
	//
	// }
}

// if curr == "]" {
//   gs.grid[new_pos.y][new_pos.x-1] = "["
//   gs.grid[new_pos.y][new_pos.x] = "]"
//   gs.grid[pos.y][pos.x] = "."
//   gs.grid[pos.y][pos.x-1] = "."
// }
// if curr == "[" {
//   gs.grid[new_pos.y][new_pos.x] = "["
//   gs.grid[new_pos.y][new_pos.x+1] = "]"
//   gs.grid[pos.y][pos.x] = "."
//   gs.grid[pos.y][pos.x+1] = "."
// }

// func (gs * GameSettings) navigate() {
//   m := map[string][2]int{
//     "v": {0, 1},
//     "<": {-1, 0},
//     "^": {0, -1},
//     ">": {1, 0},
//   }
//   for _, val := range gs.dirs {
//     dir, ok := m[val]
//     if !ok {
//       panic("Should exist")
//     }
//     new_pos := getNewPos(gs.pos, dir)
//     obj := gs.grid[new_pos.y][new_pos.x]
//     if obj == "#" {
//       continue;
//     } else if obj == "O" && !gs.moveBox(new_pos, dir){
//         continue;
//     }
//     gs.grid[gs.pos.y][gs.pos.x] = "."
//     gs.pos = new_pos
//     gs.grid[gs.pos.y][gs.pos.x] = "@"
//   }
// }

// func (gs *GameSettings) moveBox(pos Position, dir [2]int) bool {
//   new_pos := getNewPos(pos, dir)
//   obj := gs.grid[new_pos.y][new_pos.x]
//   if obj == "." {
//     gs.grid[new_pos.y][new_pos.x] = "O"
//     gs.grid[pos.y][pos.x] = "."
//     return true
//   } else if obj == "#" {
//     return false
//   }
//
//   canMove :=  gs.moveBox(new_pos, dir)
//   if canMove {
//     gs.grid[new_pos.y][new_pos.x] = "O"
//     gs.grid[pos.y][pos.x] = "."
//   }
//   return canMove;
// }
//
// func (gs *GameSettings) calcValue() int {
//   total := 0
//   for y, row := range gs.grid {
//     for x, val := range row {
//       if val == "O" {
//         total += getCoord(Position{x, y})
//       }
//     }
//   }
//   return total
// }

func getCoord(pos Position) int {
	return 100*pos.y + pos.x
}
