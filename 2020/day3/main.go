package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// ..##.......
// #...#...#..
// .#....#..#.
// ..#.#...#.#
// .#...##..#.
// ..#.##.....
// .#.#.#....#
// .#........#
// #.##...#...
// #...##....#
// .#..#...#.#

func main() {
	strs := readFile()
  
  var grid [][]string = make([][]string, len(strs))

  for idx, line := range lines {
    grid[idx] = strings.Split(line, "")
  }
	cols := [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 1}}
	for idx, val := range strs {
		if idx == 0 {
			continue
		}
		for i, colp := range cols {
			col[i] = (col[i] + colp) % len(val)
			if string(val[col[i]]) == "#" {
				trees[i] += 1
			}
		}
	}
  mult := 1
  for _, val := range trees {
    mult*= val
  }
	fmt.Println(mult*25)
}

func readFile() []string {
	bts, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Failed to read file")
	}

	cleanStr := strings.TrimRight(string(bts), "\n")
	var str []string
	for _, val := range strings.Split(cleanStr, "\n") {
		str = append(str, val)
	}
	return str
}
