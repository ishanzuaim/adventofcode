package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
    log.Fatal("Failed to convert: ", err)
	}
	return i
}

func convAll(str []string) []int {
  nums := make([]int, len(str))
  for idx, val := range str {
    nums[idx] = convToInt(val)
  }
  return nums
}

func convToStr(s int) string {
	return strconv.Itoa(s)
}

func readFile() []string {
	bts, err := os.ReadFile("inp.txt")
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


type Position struct {
	x, y int
}

func (p Position) equals(pos Position) bool {
  return p.x == pos.x && p.y == pos.y
}

func (p Position) string() string {
	return fmt.Sprintf("%d-%d", p.x, p.y)
}

