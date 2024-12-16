package main

import (
	"fmt"
	"math"
	"regexp"
)

type GameSettings struct {
	a     State
	b     State
	prize State
}

type State struct {
	x int
	y int
}

func getState(a []string) State {
	return State{
		convToInt(a[0]),
		convToInt(a[1]),
	}
}

// Button A: X+94, Y+34
// Button B: X+22, Y+67
// Prize: X=8400, Y=5400

//A (94x, 34y)
//B (22x, 67y)
//94a + 22b = 8400x
//34a + 67b = 5400y
//

//(94*67)x + (22*67)y = (8400*67)
//(34*22)x + (67*22)y = (5400*22)
// 94a = 34b

func main() {
  scale := 10000000000000
	strs := readFile()
	rp := regexp.MustCompile(`\d+`)
	total := float64(0)
	for i := 0; i < len(strs); i += 4 {
		game := strs[i : i+3]
		a := rp.FindAllString(game[0], -1)
		b := rp.FindAllString(game[1], -1)
		prize := rp.FindAllString(game[2], -1)
		gs := &GameSettings{getState(a), getState(b), State{convToInt(prize[0]) + scale , convToInt(prize[1]) + scale}}
		val := gs.simult()
		total += val
	}
	// __AUTO_GENERATED_PRINT_VAR_START__
	fmt.Println(fmt.Sprintf("main total: %v", total)) // __AUTO_GENERATED_PRINT_VAR_END__
}

func (gs *GameSettings) simult() float64 {
	a := gs.a.x * gs.b.y      //94 * 67
	ap := gs.prize.x * gs.b.y //8400  * 67
	b := gs.b.x * gs.a.y      //34 * 22
	bp := gs.prize.y * gs.b.x //5400 * 22
	var valX float64
	var valY float64
	valX = float64(ap-bp) / float64(a-b)
	if valX != math.Trunc(valX)  {
		return 0
	}
	valY = (float64(gs.prize.x) - float64(float64(gs.a.x)*valX)) / float64(gs.b.x) //94*valX + 22y = 8400
	if valY != math.Trunc(valY)  {
		return 0
	}
	return valX*3 + valY
}

// func (gs *GameSettings) calcUptoPrize(st State, count float64) float64 {
//   // __AUTO_GENERATED_PRINT_VAR_START__
//   fmt.Println(fmt.Sprintf("calcUptoPrize count: %v", count)) // __AUTO_GENERATED_PRINT_VAR_END__
//   // __AUTO_GENERATED_PRINT_VAR_START__
//   fmt.Println(fmt.Sprintf("calcUptoPrize st: %v", st)) // __AUTO_GENERATED_PRINT_VAR_END__
//   if st.x == gs.prize.x && st.y == gs.prize.y {
//     return count;
//   }
//   if st.x > gs.prize.x || st.y > gs.prize.y {
//     return math.Inf(1);
//   }
//   countA := gs.calcUptoPrize(State{st.x + gs.a.x, st.y + gs.a.y}, count + 1)
//   countB := gs.calcUptoPrize(State{st.x + gs.b.x, st.y + gs.b.y}, count + 1)
//   return math.Min(countA, countB)
// }
