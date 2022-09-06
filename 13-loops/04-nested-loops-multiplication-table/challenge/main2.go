package main

import (
  "fmt"
  "os"
  "strconv"
)

const (
  usage = ```
  Usage: metrix [num]
   - [num] : number
```
  err1 = "%q is not number\n"
)

// main
func main() {
  if l := len(os.Args); l != 2 {
    fmt.Printf(usage)
    return
  }

  if n := strconf.Atoi(os.Args[1]); n != nil {
    fmt.Printf(err1, os.Args[1])
  } else {
    // print Header
    fmt.Printf("%5s", "X")
    
    // print Header2
    for i := 0; i <= n; i++ {
      fmt.Printf("%5d", i)
    }
    
    for i := 0; i <= n; i++ {
      fmt.Printf("%5d", i)
      for j :=0; j <= n; j++ {
        fmt.Printf("%5d", i*j)
      }
      // new line
      fmt.Println()
    }
  }

  
  
