package main

import (
	// "bufio"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type State struct {
	x int
	y int
}

type Robot struct {
	Position State
	Velocity State
}

func main() {
	strs := readFile()

	re := regexp.MustCompile(`-?\d+`)
	f, err := os.Create("/tmp/yourfile")
	if err != nil {
		panic("asd")
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	for i := range 10000 {
    // __AUTO_GENERATED_PRINT_VAR_START__
    fmt.Println(fmt.Sprintf("main i: %v", i)) // __AUTO_GENERATED_PRINT_VAR_END__
    fmt.Fprintf(w, "%d seconds\n", i)
		states := make([]State, 0)
		for _, val := range strs {
			matches := re.FindAllString(val, -1)
			nums := convAll(matches)
			rb := &Robot{
				State{nums[0], nums[1]},
				State{nums[2], nums[3]},
			}
			st := rb.posAfterN(i)
			states = append(states, st)
		}
		visualizeState(w, states)
	}
}

var x int = 101
var y int = 103

func visualizeState(w *bufio.Writer, states []State) {
	for idxY := range y {
		for idxX := range x {
			if stateContains(State{idxX, idxY}, states) {
				fmt.Fprintf(w, "O")
			} else {
				fmt.Fprintf(w, ".")
			}
		}
		fmt.Fprintf(w, "\n")
	}
}

func stateContains(st State, states []State) bool {
	for _, val := range states {
		if val.x == st.x && val.y == st.y {
			return true
		}
	}
	return false
}

func getMultiple(quads []int) int {
	mult := 1
	for _, val := range quads {
		if val != 0 {
			mult *= val
		}
	}
	return mult
}

func putInQuad(st State, quads []int) {
	localY := -1
	if st.x < x/2 {
		if st.y < y/2 {
			localY = 0
		} else if st.y > y/2 {
			localY = 2
		}
	} else if st.x > x/2 {
		if st.y < y/2 {
			localY = 1
		} else if st.y > y/2 {
			localY = 3
		}
	}
	if localY == -1 {
		return
	}
	quads[localY] += 1
}

func (rb *Robot) posAfterN(n int) State {

	curr := rb.Position
	for range n {
		curr.x += rb.Velocity.x
		curr.y += rb.Velocity.y
		if curr.x < 0 {
			curr.x = x + curr.x
		} else {
			curr.x %= x
		}
		if curr.y < 0 {
			curr.y = y + curr.y
		} else {
			curr.y %= y
		}
	}
	return curr
}
