package main

import (
	"fmt"
	"slices"
	"strings"
)

var degree = map[string][]int{
	"90":   {0, -1},
	"180":  {1, 0},
	"270": {0, 1},
	"0": {-1, 0},
}

type Position struct {
	x int
	y int
}

type GameSettings struct {
	startPos Position
	grid     [][]string
}

func main() {
	strs := readFile()
	var grid [][]string = make([][]string, len(strs))
	for idx, val := range strs {
		grid[idx] = strings.Split(val, "")
	}
	x, y := getStartIdx(grid)
	startPos := Position{x, y}
	gs := &GameSettings{startPos, grid}
	pos := 0
	for _, v := range gs.getVisitedWithoutObst() {
  // for y, row := range gs.grid {
  //   for x := range row {
      // v := Position{x, y}
      if v.x == startPos.x && v.y == startPos.y {
        continue
      }
      isLoop := gs.checkIsLoop(v)
      if isLoop {
        pos += 1
      }
    }
  // }
// }
// __AUTO_GENERATED_PRINT_VAR_START__
fmt.Println(fmt.Sprintf("main pos: %v", pos)) // __AUTO_GENERATED_PRINT_VAR_END__
}

type Visited struct {
  pos   Position
  angle string
}

func (gs *GameSettings) checkIsLoop(obstactle Position) bool {
  var visited []Visited
  angle := "90"
  x, y := gs.startPos.x, gs.startPos.y
  for y+degree[angle][1] < len(gs.grid) && x+degree[angle][0] < len(gs.grid[0]) && x+degree[angle][0] >= 0 && y+degree[angle][1] >= 0 {
    newX, newY := x+degree[angle][0], y+degree[angle][1]
    nextPos := gs.grid[newY][newX]
    if nextPos == "#" || (newX == obstactle.x && newY == obstactle.y) {
      angle = turnRightAngle(angle)
      continue;
    } 
    newVisited := Visited{pos: Position{x, y}, angle: angle}
    if isVisitedBefore(visited, newVisited) {
      return true
    }
    visited = append(visited, newVisited)
    x = newX
    y = newY
    if gs.grid[y][x] == "#" {
      panic("WHATG")
    }

  }
  return false
}

func isVisitedBefore(visited []Visited, visit Visited) bool {
  for _, val := range visited {
    if val.angle == visit.angle && val.pos.x == visit.pos.x && val.pos.y == visit.pos.y {
      return true
    }
  }
  return false
}

func hasPosition(positions []Position, position Position) bool {
  for _, val := range positions {
    if position.x == val.x && position.y == val.y {
      return true
    }
  }
  return false
}

func (gs *GameSettings) getVisitedWithoutObst() []Position {
  var visited []Position

  angle := "90"
  x, y := gs.startPos.x, gs.startPos.y
  for y+degree[angle][1] < len(gs.grid) && x+degree[angle][0] < len(gs.grid[0]) && x+degree[angle][0] >= 0 && y+degree[angle][1] >= 0 {
    newX, newY := x+degree[angle][0], y+degree[angle][1]
    nextPos := gs.grid[newY][newX]
    if nextPos == "#" {
      angle = turnRightAngle(angle)
    } else {
      newVisited := Position{x, y}
      if(!hasPosition(visited, newVisited)) {
        visited = append(visited, newVisited)
      }
      x = newX
      y = newY
      if gs.grid[y][x] == "#" {
        panic("WHATG")
      }
    }
  }
  visited = append(visited, Position{x,y})
  return visited
}

func p1(strs []string) int {
  var visited [][]string = make([][]string, len(strs))
  var grid [][]string = make([][]string, len(strs))
  for idx, val := range strs {
    grid[idx] = strings.Split(val, "")
    visited[idx] = strings.Split(val, "")
  }
  x, y := getStartIdx(grid)
  visited[y][x] = "X"
  moveGuardTillGoal(x, y, grid, visited)
  total := 0
  for _, val := range visited {
    for _, item := range val {
      if item == "X" {
        total += 1
      }
    }
  }
  return total
}

func moveGuardTillGoal(x, y int, grid [][]string, visited [][]string) int {
  dx, dy := 0, -1
  total := 1
  for y+dy < len(grid) && x+dx < len(grid[0]) && x+dx >= 0 && y+dy >= 0 {
    nextPos := grid[y+dy][x+dx]
    if nextPos == "#" {
      dx, dy = turnRight(dx, dy)
    } else {
      total += 1
      x += dx
      y += dy
      visited[y][x] = "X"
    }
  }
  return total
}

func turnRightAngle(angle string) string {
  if angle == "0" {
    return "90"
  }
  if angle == "90" {
    return "180"
  }
  if angle == "180" {
    return "270"
  }
  if angle == "270" {
    return "0"
  }
  panic("what anlge")
}

func turnRight(dx, dy int) (int, int) {
  if dx == 0 && dy == -1 {
    dx, dy = 1, 0
  } else if dx == 1 && dy == 0 {
    dx, dy = 0, 1
  } else if dx == 0 && dy == 1 {
    dx, dy = -1, 0
  } else if dx == -1 && dy == 0 {
    dx, dy = 0, -1
  }
  return dx, dy
}

func getStartIdx(grid [][]string) (int, int) {
  for y, val := range grid {
    x := slices.Index(val, "^")
    if x != -1 {
      return x, y
    }
  }
  panic("Should have found something")
}
