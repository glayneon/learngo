package main

import (
  "fmt"
)


const max = 5

// main

func main() {
  // Header
  fmt.Printf("%5s", "X")
  for i := 0 ; i <= max; i++ {
    fmt.Printf("%5d", i)
  }
  
  for i := 0; i <= max; i++ {
    fmt.Printf("%5d", i)
    
    for j := 0; j <= max; j++ {
      fmt.Printf("%5d", i*j)
    }
    fmt.Println()
  }
}
