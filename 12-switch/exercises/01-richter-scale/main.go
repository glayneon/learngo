package main

// Richter Scale Program
import (
  "fmt"
  "os"
  "strconv"
)

// constants
const (
  err1 = "Give me the magnitude of the earthquake.\n"
  err2 = "I couldn't get that, sorry.\n"
  pr1 = "%3.3f is %s\n"
)


func main() {
  if len(os.Args) != 2 {
    fmt.Printf(err1)
    return
  }
  
  a1 := os.Args[1]
  if n, err := strconv.ParseFloat(a1, 32); err != nil {
    fmt.Printf(err2)
    return
  } else {
    switch {
      case n >= 10:
      fmt.Printf(pr1, n, "massive")
      case n >= 8:
      fmt.Printf(pr1, n, "great")
      case n >= 7:
      fmt.Printf(pr1, n, "major")
      case n >= 6:
      fmt.Printf(pr1, n, "strong")
      case n >= 5:
      fmt.Printf(pr1, n, "moderate")
      case n >= 4:
      fmt.Printf(pr1, n, "light")
      case n >= 3:
      fmt.Printf(pr1, n, "minor")
      case n >= 2:
      fmt.Printf(pr1, n, "very minor")
      default:
      fmt.Printf(pr1, n, "micro")
    }
  }
}
