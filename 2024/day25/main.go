package main

import "fmt"

func main() {
	strs := readFile()
  var locks [][5]int
  var keys [][5]int

	for idx := range strs {
		if idx%8 == 0 {
      isLock, height := processGrid(strs[idx+1 : idx+8])
      if isLock {
        locks = append(locks, height)
      } else {
        keys = append(keys, height)
      }
		}
	}


  total := 0
  for _, lock := range locks {
    for _, key := range keys {
      if isCompatible(key, lock) {
        total += 1
      }
    }
  }
  // __AUTO_GENERATED_PRINT_VAR_START__
  fmt.Println(fmt.Sprintf("main total: %v", total)) // __AUTO_GENERATED_PRINT_VAR_END__
}

func isCompatible(key, lock [5]int) bool {
  for idx := range 5 {
    if key[idx] + lock[idx] > 5 {
      return false
    }
  }
  return true
}

func processGrid(str []string) (bool, [5]int) {
	var height [5]int
	isLock := str[0] == "#####"

	for col := range 5 {
		for row := range 6 {
			if row == 0 {
				continue
			}
			if isLock && str[row][col] == '#' {
				height[col]++
			} else if !isLock && str[row][col] == '.' {
				height[col]++
			} else {
				break
			}
		}
	}
  if !isLock {
    height = inverse(height)
  }
  return isLock, height
}

func inverse(nt [5]int) [5]int {
	var height [5]int
  for idx, val := range nt {
    height[idx] =  5 - val
  }
  return height
}
