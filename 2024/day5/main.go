package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func main() {
  strs := readFile()
  var input []string;
  var expected [][]string
  flag := 0
  for _, val := range strs {
    if len(val) == 0 {
      flag = 1;
      continue;
    }
    if(flag == 0) {
      input = append(input, val)
    } else {
      expected = append(expected, strings.Split(val, ","))
    }
  }

  out := generateMap(input)

  sum := 0
  for _, expt := range expected {
      if !isPrintCorrect(expt, out) {
        fixed := fixPrint(expt, out)
        sum += getMiddle(fixed)
    }
  }
  // __AUTO_GENERATED_PRINT_VAR_START__
  fmt.Println(fmt.Sprintf("main sum: %v", sum)) // __AUTO_GENERATED_PRINT_VAR_END__
}

func isIndexCorrect(idx int, print []string,out map[string][]string) bool {
  before := print[0:idx]
  return checkIsCorrect(before, out[string(print[idx])])
}

func isPrintCorrect(print []string, out map[string][]string) bool {
    for idx := range print {
      if !isIndexCorrect(idx, print, out) {
      return false
    }
  }
  return true
}

func fixPrint(print []string, out map[string][]string) []string {
  for i := 0; i < len(print); i++ {
    if !isIndexCorrect(i, print, out) {
      if i == 0 {
        panic("should be correct")
      }
      print[i -1], print[i] = print[i], print[i -1]
    }
  }
  
  if(isPrintCorrect(print,out)) {
    return print
  }
  return fixPrint(print, out)
}


func getMiddle(ls []string) (int) {
    half := len(ls) / 2
    return convToInt(ls[half])
}

func part1(expected [][]string, out map[string][]string)(int) {
  validIdx := []int{}
  for ei, expt := range expected {
    flag := true
    for idx, item := range expt {
      before := expt[0:idx]
      if !checkIsCorrect(before, out[string(item)]) {
        flag = false
      }
    }
    if flag {
      validIdx = append(validIdx, ei)
    }
  }
  sum := 0
  for _, idx := range validIdx {
    exp := expected[idx]
    half := len(exp) / 2
    sum += convToInt(exp[half])
  }
  return sum
}

func generateMap(inp []string) (map[string][]string) {
  m := make(map[string][]string)

  for _, item := range inp {
    matched :=regexp.MustCompile(`\d+`).FindAllString(item, -1)
    _, ok := m[matched[0]]
    if !ok {
      m[matched[0]] = []string{matched[1]}
    } else {
      m[matched[0]] = append(m[matched[0]], matched[1])
    }
  }
  return m
}

func checkIsCorrect(before, blocked []string) bool {
  for _, val := range before {
    ok := slices.Contains(blocked, val)
    if ok {
      return false;
    }
  }
  return true;
}
