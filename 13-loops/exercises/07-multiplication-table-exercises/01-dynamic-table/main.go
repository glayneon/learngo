package main

import (
  "fmt"
  "strconv"
  "os"
)

const (
  usage = ```
  Give me the size of the table\n
  ```
  
  err1 = "Wrong size\n"

func main() {
  
  if len(os.Args) != 2 {
    fmt.Printf(usage)
    return
  }
  
  // check it could be converted to integer
  if t := strconv.Atoi(os.Args[1]); t != nil || t < 0 {
    fmt.Printf(err1)
  } else {
    // Print Headers
    fmt.Printf("%5s", "X")
    
    // Print Headers
    for i := 0; i <= t; i++ {
      fmt.Printf("%5d", i)
    }
    
    // Print contents of the table
    for i := 0; i <= t; i++ {
      fmt.Printf("%5d", i)
      for j := 0; j <= t; j++ {
        fmt.Printf("%5d", i*j)
      }
      // new line
      fmt.Println()
    }
}
