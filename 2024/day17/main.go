package main

import (
	"fmt"
	"math"
	"regexp"
	"sync"
)

type GameSettings struct {
	registers []int
	outputs   []int
}

var wg sync.WaitGroup

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func main() {
	defer wg.Done()
	strs := readFile()
	re := regexp.MustCompile(`\d+`)
	B := re.FindAllString(strs[1], -1)[0]
	C := re.FindAllString(strs[2], -1)[0]
	program := convAll(re.FindAllString(strs[4], -1))

	arr := makeRange(0, 1000)
	ch := make(chan int)
	for _, val := range arr {
		wg.Add(1)
		go func(r int) {
			i := r * 200_000_000
			for {
				gs := &GameSettings{[]int{i, convToInt(B), convToInt(C)}, make([]int, 0)}
				ctr := 0
				for ctr < len(program)-1 {
					ret := gs.processOp(program[ctr], program[ctr+1])
					if ret != -1 {
						ctr = ret
					} else {
						ctr += 2
					}
				}
				if isEqual(gs.outputs, program) {
					panic(i)
				}
				i += 1
				if i > ((r + 1) * 200_000_000) {
					panic("Asd")
				}
			}
		}(val)
	}

	data := <-ch
	// __AUTO_GENERATED_PRINT_VAR_START__
	fmt.Println(fmt.Sprintf("main data: %v", data)) // __AUTO_GENERATED_PRINT_VAR_END__
	close(ch)
}

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, val := range a {
		if b[idx] != val {
			return false
		}
	}
	return true
}

func (gs *GameSettings) getCombo(idx int) int {
	if idx >= 0 && idx <= 3 {
		return idx
	}
	return gs.registers[idx%4]
}

func (gs *GameSettings) processOp(opc, ope int) int {
	if opc == 0 {
		combo := gs.getCombo(ope)
		gs.registers[0] = gs.registers[0] / int(math.Pow(2, float64(combo)))
	}
	if opc == 1 {
		gs.registers[1] = gs.registers[1] ^ ope
	}
	if opc == 2 {
		combo := gs.getCombo(ope)
		gs.registers[1] = combo % 8
	}
	if opc == 3 && gs.registers[0] != 0 {
		return ope
	}
	if opc == 4 {
		gs.registers[1] = gs.registers[1] ^ gs.registers[2]
	}
	if opc == 5 {
		combo := gs.getCombo(ope)
		gs.outputs = append(gs.outputs, combo%8)
	}
	if opc == 6 {
		combo := gs.getCombo(ope)
		gs.registers[1] = gs.registers[0] / int(math.Pow(2, float64(combo)))
	}
	if opc == 7 {
		combo := gs.getCombo(ope)
		gs.registers[2] = gs.registers[0] / int(math.Pow(2, float64(combo)))
	}
	return -1
}
