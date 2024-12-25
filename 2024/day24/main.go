package main

import (
	"fmt"
	"strconv"
	"strings"
)

type GameSettings struct {
	wires map[string]int
	eqns  [][4]string
	max   int
}

func main() {
	wires := make(map[string]int)
	eqns := make([][4]string, 0)
	strs := readFile()
	flag := false
	maxV := 0
	for _, val := range strs {
		if val == "" {
			flag = true
			continue
		}
		if !flag {
			vals := strings.Split(val, ": ")
			wires[vals[0]] = convToInt(vals[1])
			if strings.HasPrefix(val, "z") {
				aft, _ := strings.CutPrefix(val, "z")
				num := convToInt(aft)
				if maxV < num {
					maxV = num
				}
			}
		} else {
			vals := strings.Split(val, " ")
			if strings.HasPrefix(vals[4], "z") {
				aft, _ := strings.CutPrefix(vals[4], "z")
				num := convToInt(aft)
				if maxV < num {
					maxV = num
				}
			}
			eqns = append(eqns, [4]string{vals[0], vals[1], vals[2], vals[4]})
		}
	}
	gs := &GameSettings{wires, eqns, maxV}
  for len(gs.eqns) > 0 {
    gs.loop()
  }
	bte := make([]string, maxV+1)
	for k, v := range wires {
		if strings.HasPrefix(k, "z") {
			aft, _ := strings.CutPrefix(k, "z")
			num := convToInt(aft)
			bte[maxV-num] = convToStr(v)
		}
	}
  val := strings.Join(bte, "")
  ans, _ := strconv.ParseInt(val, 2, 64)
	// __AUTO_GENERATED_PRINT_VAR_START__
	fmt.Println(fmt.Sprintf("main bte: %v", ans)) // __AUTO_GENERATED_PRINT_VAR_END__
}

func (gs *GameSettings) loop() {
	new_eqn := make([][4]string, 0)
	for _, val := range gs.eqns {
		valA, checkA := gs.wires[val[0]]
		valB, checkB := gs.wires[val[2]]
		if !checkA || !checkB {
			new_eqn = append(new_eqn, val)
			continue
		}
		gs.wires[val[3]] = getResult(valA, valB, val[1])
	}
	gs.eqns = new_eqn
}

func getResult(a, b int, eqn string) int {
	if eqn == "XOR" {
		return a ^ b
	}
	if eqn == "OR" {
		return a | b
	}
	if eqn == "AND" {
		return a & b
	}
	panic("NO")
}
