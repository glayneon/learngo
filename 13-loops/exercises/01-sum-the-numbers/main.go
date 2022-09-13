package main

import (
  "fmt"
)

const (
  max = 10
)

func main() {
  sum := 0
  for i := 0; i <= max; i++ {
    sum += i
  }
  
  fmt.Printf("Sum: %4d", sum)
}
