package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

// message
const (
  pr1 = "%q has %d days.\n"
  err1 = "Give me a month name.\n"
  err2 = "%q is not a month.\n"
  now = time.Now().Year()
  
  // month
  jan = "january"
  feb = "february"
  mar = "march"
  apr = "april"
  may = "may"
  jun = "june"
  jul = "july"
  aug = "august"
  sep = "september"
  oct = "october"
  nov = "november"
  dec = "december"
)

// main
func main() {
  a := os.Args
  
  // chech arguments is two.
  if len(a) != 2 {
    fmt.Printf(err1)
    return
  }
  
  // check month
  m := strings.ToLower(a[1])
  
  if m == jan {
    fmt.Printf(pr1, m, 31)
  } else if m == feb {
    if now % 4 == 0 {
      fmt.Printf(pr1, m, 29)
    } else {
      fmt.Printf(pr1, m, 28)
    }
  } else if m == mar {
    fmt.Printf(pr1, m, 31)
  } else if m == apr {
    fmt.Printf(pr1, m, 30)
  } else if m == may {
    fmt.Printf(pr1, m, 31)
  } else if m == jun {
    fmt.Printf(pr1, m, 30)
  } else if m == jul {
    fmt.Printf(pr1, m, 31)
  } else if m == aug {
    fmt.Printf(pr1, m, 31)
  } else if m == sep {
    fmt.Printf(pr1, m, 30)
  } else if m == oct {
    fmt.Printf(pr1, m, 31)
  } else if m == nov {
    fmt.Printf(pr1, m, 30)
  } else if m == dec {
    fmt.Printf(pr1, m, 31)
  } else {
    fmt.Printf(err2, m)
  }
}
