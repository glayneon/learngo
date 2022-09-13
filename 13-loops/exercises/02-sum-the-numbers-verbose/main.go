package main

import (
  "fmt"
)

const (
  max = 10
)

func main() {
  sum := 0
  for i := 1; i <= 10; i++ {
    sum :+ i
    if i != 10 {
      fmt.Print(i, "+")
    } else {
      fmt.Print(i, "=")
    }
  }
  fmt.Printf("%3d", sum)
}
