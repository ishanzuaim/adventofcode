package main

import (
	"fmt"
	"strings"
)

type GameSettings struct {
	inputs []string
  inpMap map[string]int
}

func main() {
	strs := readFile()
	inputs := strings.Split(strs[0], ", ")
	gs := GameSettings{inputs, make(map[string]int)}
	total := 0

	for _, val := range strs[2:] {
		total += gs.isPossible(val)
	}
	// __AUTO_GENERATED_PRINT_VAR_START__
	fmt.Println(fmt.Sprintf("main total: %v", total)) // __AUTO_GENERATED_PRINT_VAR_END__
}

func (gs GameSettings) isPossible(str string) int {
	if str == "" {
		return 1
	}
	total := 0
	for _, val := range gs.inputs {
		if strings.HasPrefix(str, val) {
      newVal := strings.TrimPrefix(str, val)
      // __AUTO_GENERATED_PRINT_VAR_START__
      fmt.Println(fmt.Sprintf("isPossible newVal: %v", newVal)) // __AUTO_GENERATED_PRINT_VAR_END__
      amt, ok := gs.inpMap[newVal]
      if !ok {
        amt = gs.isPossible(newVal)
        gs.inpMap[newVal] = amt
      }
			total += amt
		}
	}
	return total
}
