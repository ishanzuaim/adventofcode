package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func countStr(pwd string, chr rune) int {
	cnt := 0
	for _, ichr := range pwd {
		if ichr == chr {
			cnt++
		}
	}
	return cnt
}

func isValid(pwd string, chr string, pos1, pos2 int) bool {
  count := 0
  if string(pwd[pos1 - 1 ]) == chr {
    count ++;
  }
  if string(pwd[pos2 - 1]) == chr {
    count ++;
  }
  return count == 1
}

func main() {
	strs := readFile()
	valid := 0
	for _, val := range strs {
    var lower, upper int
    var char, pwd string

    _, err := fmt.Sscanf(val, "%d-%d %1s: %s", &lower, &upper, &char, &pwd);
    if err != nil {
      log.Fatal("Failed to scan: ", err.Error())
    }
		if isValid(pwd, char, lower, upper) {
			valid++
		}
	}
	fmt.Println(valid)
}

func convToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Failed to convert")
	}
	return i
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
