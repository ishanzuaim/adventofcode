package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
  file, err := os.Open("./input.txt")
  if err != nil {
    log.Fatal("Failed to open file")
  }
  
  defer file.Close()

  scanner := bufio.NewScanner(file);

  var s []int;
  for scanner.Scan() {
    val, err := strconv.Atoi(scanner.Text())
    if err != nil {
      log.Fatal("Failed to conv string")
    }
    s = append(s, val)
  }
  
  for i := 0; i < len(s); i++ {
    for j := i; j < len(s); j++ {
      for k := j; k < len(s); k++ {
      if s[j] + s[i] + s[k]  == 2020 {
        fmt.Println(s[i] * s[j] * s[k])
          return;
        }
      }
    }
  }
}

