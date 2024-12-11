package main

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

type GameSettings struct {
	memo map[int]map[int]int
}

func main() {
	var wg sync.WaitGroup
	strs := readFile()
	nums := convAll(strings.Split(strs[0], " "))
	results := make(chan int, len(nums))


	for _, val := range nums {
		wg.Add(1)
		// sum += processValue(nums[0], 0)
		go func(num int) {
      memo := make(map[int]map[int]int)
      gs := &GameSettings{memo}
			defer wg.Done()
			result := gs.processValue(num, 0)
			// __AUTO_GENERATED_PRINTF_START__
			fmt.Println("main 1") // __AUTO_GENERATED_PRINTF_END__
			results <- result
		}(val)
	}

	wg.Wait()
	close(results)
	sum := 0
	for val := range results {
		sum += val
	}
	// __AUTO_GENERATED_PRINT_VAR_START__
	fmt.Println(fmt.Sprintf("main sum: %v", sum)) // __AUTO_GENERATED_PRINT_VAR_END__
}

func (gs *GameSettings) processValue(num, length int) int {
	if length == 75 {
		return 1
	}
	if num == 0 {
		return gs.memoizedProcessValue(1, length+1)
	}
	size := intSize(num)
	if size%2 == 0 {
		a, b := splitInt(num, size)
		return gs.memoizedProcessValue(a, length+1) + gs.memoizedProcessValue(b, length+1)
	} else {
		return gs.memoizedProcessValue(num*2024, length+1)
	}
}

func (gs *GameSettings) memoizedProcessValue(num, length int) int {
	_, ok := gs.memo[num]
	if !ok {
		gs.memo[num] = make(map[int]int)
	}
	val, ok := gs.memo[num][length]
	if ok {
		return val
	}
	val = gs.processValue(num, length)
	gs.memo[num][length] = val
	return val
}

func intSize(num int) int {
	count := 0
	for num > 0 {
		num = num / 10
		count++
	}
	return count
}

func splitInt(num, size int) (int, int) {
	ten := int(math.Pow10(size / 2))
	b := num / ten
	a := num % ten
	return a, b
}

func splitString(str string) []int {
	half := len(str) / 2
	return []int{convToInt(str[0:half]), convToInt(str[half:])}
}

func p1(nums []int) {

	for i := 0; i < 75; i++ {
		// __AUTO_GENERATED_PRINT_VAR_START__
		fmt.Println(fmt.Sprintf("main blink: %v, len: %v", i, len(nums))) // __AUTO_GENERATED_PRINT_VAR_END__
		new_nums := []int{}
		for _, val := range nums {
			if val == 0 {
				new_nums = append(new_nums, 1)
			} else if len(convToStr(val))%2 == 0 {
			} else {
				new_nums = append(new_nums, splitString(convToStr(val))...)
				new_nums = append(new_nums, val*2024)
			}
		}
		nums = new_nums
	}
}
