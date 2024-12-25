package main

import (
	"fmt"
	"strings"
)

type GameSettings struct {
	grid  [][]string
	start Position
}

func main() {
	strs := readFile()
	gs := &GameSettings{
		make([][]string, 73),
		Position{-1, -1},
	}
	for y := range 73 {
		gs.grid[y] = make([]string, 73)
		if y == 0 || y == 72 {
			for x := range 73 {
				gs.grid[y][x] = "#"
			}
		} else {
			for x := range 73 {
				if x == 0 || x == 72 {
					gs.grid[y][x] = "#"
				} else {
					gs.grid[y][x] = "."
				}
			}
		}
	}

	fmt.Println(strs[0:2930])
	for _, val := range strs[0:2929] {
		nums := convAll(strings.Split(val, ","))
		gs.grid[nums[1]+1][nums[0]+1] = "#"
	}
	gs.start = Position{1, 1}
	val := gs.bfs()
	// __AUTO_GENERATED_PRINT_VAR_START__
	fmt.Println(fmt.Sprintf("main val: %v", val)) // __AUTO_GENERATED_PRINT_VAR_END__
}

func findMin(queue []State) (int, State) {
	min := 0
	idx := -1
	for x, val := range queue {
		if idx == -1 {
			min = val.cost
			idx = x
		} else if val.cost < min {
			min = val.cost
			idx = x
		}
	}
	return idx, queue[idx]
}

type State struct {
	position Position
	cost     int
}

func (gs *GameSettings) bfs() int {
	queue := []State{{gs.start, 0}}
	visited := []Position{}
	for len(queue) > 0 {
		idx, curr := findMin(queue)
		queue = append(queue[:idx], queue[idx+1:]...)
		visited = append(visited, curr.position)

		if 71 == curr.position.x && 71 == curr.position.y {
			return curr.cost
		}

		states := gs.findSurrounding(curr)
		for _, val := range states {
			if !contains(visited, val) {
				ok, idx := hasBetterCost(queue, val)
				if ok {
					if idx != -1 {
						queue[idx] = val
					} else {
						queue = append(queue, val)
					}
				}
			}
		}
	}
	return 0
}

func hasBetterCost(queue []State, state State) (bool, int) {
	for idx, val := range queue {
		if state.position.equals(val.position) {
			if state.cost < val.cost {
				return true, idx
			} else {
				return false, -1
			}
		}
	}
	return true, -1
}

func contains(visited []Position, state State) bool {
	for _, val := range visited {
		if val.equals(state.position) {
			return true
		}
	}
	return false
}

func (gs *GameSettings) findSurrounding(st State) []State {
	states := make([]State, 0)

	pos := st.position

	//left
	checkPos := Position{pos.x - 1, pos.y}
	if gs.getVal(checkPos) != "#" {
		states = append(states, State{checkPos, st.cost + 1})
	}

	//right
	checkPos = Position{pos.x + 1, pos.y}
	if gs.getVal(checkPos) != "#" {
		states = append(states, State{checkPos, st.cost + 1})
	}

	checkPos = Position{pos.x, pos.y - 1}
	if gs.getVal(checkPos) != "#" {
		states = append(states, State{checkPos, 1 + st.cost})
	}

	checkPos = Position{pos.x, pos.y + 1}
	if gs.getVal(checkPos) != "#" {
		states = append(states, State{checkPos, 1 + st.cost})
	}

	return states
}

func (gs *GameSettings) findVal(str string) Position {
	for y, row := range gs.grid {
		for x, val := range row {
			if val == str {
				return Position{x, y}
			}
		}
	}
	return Position{-1, -1}
}

func (gs *GameSettings) getVal(pos Position) string {
	return gs.grid[pos.y][pos.x]
}

var angles []string = []string{"90", "180", "270", "0"}

func getNextPositionFromAngle(pos Position, angle int) Position {
	if angle == 0 {
		return Position{pos.x + 1, pos.y}
	}
	if angle == 1 {
		return Position{pos.x, pos.y + 1}
	}
	if angle == 2 {
		return Position{pos.x - 1, pos.y}
	}
	if angle == 3 {
		return Position{pos.x, pos.y - 1}
	}
	panic("WHAT")
}

// func (gs *GameSettings) traverse(pos Position, angle, cost int) (int, []Visited) {
// 	v := Visited{pos, angle, cost}
// 	curr := gs.getVal(pos)
// 	if curr == "#" {
// 		return int(math.Inf(0)), nil
// 	} else if curr == "E" {
// 		return cost, nil
// 	}
//   idx :=  contains(gs.visited, v)  //nogood
//   if idx != -1 {
//     if gs.visited[idx].cost < cost {
//       return int(math.Inf(0)), nil
//     }
//     gs.visited[idx] = v
//   }
//   if idx == -1 {
//     gs.visited = append(gs.visited, v) //nogood
//   }
// 	costA, _ := gs.traverse(getNextPositionFromAngle(pos, angle), angle, cost+1)
// 	costB, _ := gs.traverse(pos, (angle+1)%4, cost+1000)
// 	costC, _ := gs.traverse(pos, (angle+3)%4, cost+1000)
//
//   if costA < costB && costA < costC {
//     return costA, nil
//   }
//   if costB <= costA && costB <= costC {
//     return costB, nil
//   }
//   if costC <= costA && costC <= costB {
//     return costC, nil
//   }
//   panic("WHAT")
// }
