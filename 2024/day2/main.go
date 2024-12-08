package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
  val := readFile()
  isIncreasing := false 
  isDecreasing := false

  count := 0
  for _, row := range val {
	for i := 1; i < len(row); i++ {
		diff := convToInt(string(row[i])) - convToInt(string(row[i-1]))
		if diff > 0 {
			isIncreasing = true
		} else if diff < 0 {
			isDecreasing = true
		} else {
        continue;
      }

		if isDecreasing && isIncreasing {
        continue
		}

		if diff > 3 || diff < -3 {
        continue
		}
      count++;
	}
  }
  fmt.Println(count)
}

func convToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Failed to convert")
	}
	return i
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
