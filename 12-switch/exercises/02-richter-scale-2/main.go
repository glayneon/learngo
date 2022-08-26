package main

import (
  "os"
  "fmt"
  "strings"
)

const (
  err1 = "Tell me the magnitude of the earthquake in human terms.\n"
  err2 = "%s's richter scale is unknown\n"
  pr1 = "%s's richter scale is %s\n"
  micro = "less than 2.0\n"
  vminor = "2 - 2.9\n"
  minor = "3 - 3.9\n"
  light = "4 - 4.9\n"
  mod = "5 - 5.9\n"
  strong = "6 - 6.9\n"
  major = "7 - 7.9\n"
  great = "8 - 9.9\n"
  mass = "10+\n"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Printf(err1)
    return
  }
  
  switch n := strings.ToLower(os.Args[1]); {
  case "micro":
    fmt.Printf(pr1, n, micro)
  case "very minor":
    fmt.Printf(pr1, n, vminor)
  case "minor":
    fmt.Printf(pr1, n, minor)
  case "light":
    fmt.Printf(pr1, n, light)
  case "moderate":
    fmt.Printf(pr1, n, mod)
  case "strong":
    fmt.Printf(pr1, n, strong)
  case "major":
    fmt.Printf(pr1, n, major)
  case "great":
    fmt.Printf(pr1, n, great)
  case "massive":
    fmt.Printf(pr1, n, mass)
  default:
    fmt.Printf(err2, n)
}
