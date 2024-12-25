package main

import "fmt"

func main() {
	strs := readFile()
	// for _, val := range strs {
	// 	println(val)
	// }
	total := 0
	for _, val := range strs {
		num := convToInt(val)
		for range 2000 {
			num = calcSecret(num)
		}
		total += num
	}
	// __AUTO_GENERATED_PRINT_VAR_START__
	fmt.Println(fmt.Sprintf("main potal: %v", total)) // __AUTO_GENERATED_PRINT_VAR_END__
}

func calcSecret(initial int) int {
	valA := ((initial * 64) ^ initial) % 16777216
	valB := ((valA / 32) ^ valA) % 16777216
	valC := ((valB * 2048) ^ valB) % 16777216
	return valC
}
