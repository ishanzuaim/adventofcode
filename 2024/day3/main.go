package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str := strings.Join(readFile(), "")
  re := regexp.MustCompile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)
  matched := re.FindAllString(str, -1)
  flag := 1
  total := 0
  for _, val := range matched {
    if val == "do()" {
      flag = 1
    } else if val == "don't()" {
      flag = 0
    } else if flag == 1 {
      // mul(1, 2)
      re2 := regexp.MustCompile(`\d+`)
      nums := re2.FindAllString(val, -1)
      total += convToInt(nums[0]) * convToInt(nums[1])
    }
  }

}
