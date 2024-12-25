package main

import (
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

