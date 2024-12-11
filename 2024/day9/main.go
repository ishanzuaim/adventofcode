package main

import (
	"fmt"
	"strings"
)

type GameSettings struct {
	expandedStr []int
	countMap    map[int]int
}

func main() {
	gs := &GameSettings{}
	strs := readFile()
	gs.init(strs)
	highest := len(strs[0]) / 2
	gs.modifyGenP2(highest)
	sum := gs.calcCheckSums()
	// __AUTO_GENERATED_PRINT_VAR_START__
	fmt.Println(fmt.Sprintf("main sum: %v", sum)) // __AUTO_GENERATED_PRINT_VAR_END__
}

func (gs *GameSettings) init(strs []string) {
	expandedStr := []int{}
	countMap := make(map[int]int)
	intPtr := 0
	for idx, val := range strings.Split(strs[0], "") {
		size := -1
		if idx%2 == 0 {
			countMap[intPtr] = convToInt(val)
		}
		for range convToInt(val) {
			if idx%2 == 0 {
				expandedStr = append(expandedStr, intPtr)
			} else {
				expandedStr = append(expandedStr, size)
				size -= 1
			}
		}
		if idx%2 == 0 {
			intPtr += 1
		}
	}
	gs.expandedStr = expandedStr
	gs.countMap = countMap
}

func (gs *GameSettings) calcCheckSums() int {
	sum := 0
	// idx := 0
	for idx, val := range gs.expandedStr {
		if val > 0 {
			sum += idx * val
		}
	}
	return sum
}

func (gs *GameSettings) modifyGenP2(highest int) {
	for highest >= 0 {
		idx := 0
		size := gs.countMap[highest]
		for idx <= len(gs.expandedStr) {
			val := gs.expandedStr[idx]
			if val == highest {
				break
			}
			if gs.expandedStr[idx] == -size {
				gs.clearFromIdx(idx, highest)
				break
			}
			idx += 1
		}
		gs.tidyUp()
		highest -= 1
	}
}

func (gs *GameSettings) tidyUp() {
	incr := 0
	for idx, val := range gs.expandedStr {
		if val < 0 {
			gs.expandedStr[idx] = -1 - incr
			incr += 1
		} else {
			incr = 0
		}
	}
}

func (gs *GameSettings) clearFromIdx(idx, new_val int) {
	val := gs.expandedStr[idx]

	incr := 0
	for idx, val := range gs.expandedStr {
		if val == new_val {
			gs.expandedStr[idx] = -1 - incr
			incr += 1
		}
	}

	backIdx := idx
	for val < 0 {
		gs.expandedStr[backIdx] = new_val
		backIdx -= 1
		val = gs.expandedStr[backIdx]
	}

	forIdx := idx + 1
	incr = 0
	if forIdx >= len(gs.expandedStr) {
		return
	}
	val = gs.expandedStr[forIdx]
	for val < 0 || forIdx >= len(gs.expandedStr) {
		gs.expandedStr[forIdx] = -1 + incr
		forIdx -= 1
		incr += 1
		val = gs.expandedStr[forIdx]
	}
}

func (gs *GameSettings) findEmpty(size int) int {
	for idx, val := range gs.expandedStr {
		if val == -size {
			return idx
		}
	}
	return -1
}

func (gs *GameSettings) calcCheckSum() int {
	startPtr := 0
	endPtr := len(gs.expandedStr) - 1
	sum := 0

	for startPtr <= endPtr {
		num := gs.expandedStr[startPtr]
		if num != -1 {
			sum += startPtr * num
		} else {
			endNum := gs.expandedStr[endPtr]
			for endNum == -1 {
				endPtr -= 1
				if endPtr < startPtr {
					break
					// panic("What")
				}
				endNum = gs.expandedStr[endPtr]
			}
			if endNum == -1 {
				break
			}
			sum += startPtr * endNum
			endPtr -= 1
		}
		startPtr += 1
	}
	return sum
}

func replaceAtIndex(in string, chr string, i int) string {
	out := []rune(in)
	out[i] = []rune(chr)[0]
	return string(out)
}
