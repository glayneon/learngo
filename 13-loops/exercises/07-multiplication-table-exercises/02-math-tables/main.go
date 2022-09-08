package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

const (
  fmt1 = "%5s"
  fmt2 = "%5d"
  ops = "*/+-%"
  usage = ```
  Usage: [op=*/+-%] [size]
```
  err1 = "Size is missing or Negative number\n"
  err2 = ```
  Invalid operator.
  Valid ops one of: %s
```

func main() {
  // get args
  args := os.Args
  
  if len(args) != 3 {
    fmt.Printf(usage)
    return
  }
  
  // get ops and size
  o, s := args[1], args[2]
  
  // check operator
  if ok := strings.IndexAny(o, ops); ok != -1 {
    fmt.Printf(err2, o)
    return
  }
  
  // check size
  in, er := strconv.Atoi(s)
  if er != nil || i < 0 {
    fmt.Printf(err1)
    fmt.Printf(usage)
    return
  }
  
  // print Header
  fmt.Printf(fmt1, o)
  for i := 0; i <= in; i++ {
    fmt.Printf(fmt2, i)
  }
  
  switch o {
  case "*": {
    for i := 0; i <= in; i++ {
      fmt.Printf(fmt2, i)
      for j := 0; j <= in; j++ {
        fmt.Printf(fmt2, i*j)
      }
      fmt.Println()
    }
  case "/": {
    for i := 0; i <= in; i++ {
      fmt.Printf(fmt2, i)
      for j := 0; j <= in; j++ {
        fmt.Printf(fmt2, i/j)
      }
      fmt.Println()
    }
  }
  case "+": {
    for i := 0; i <= in; i++ {
      fmt.Printf(fmt2, i)
      for j := 0; j <= in; j++ {
        fmt.Printf(fmt2, i+j)
      }
      fmt.Println()
    }
  }
  case "-": {
    for i := 0; i <= in; i++ {
      fmt.Printf(fmt2, i)
      for j := 0; j <= in; j++ {
        fmt.Printf(fmt2, i-j)
      }
      fmt.Println()
    }
  case "%": {
    for i := 0; i <= in; i++ {
      fmt.Printf(fmt2, i)
      for j := 0; j <= in; j++ {
        fmt.Printf(fmt2, i%j)
      }
      fmt.Println()
    }
  }
}
