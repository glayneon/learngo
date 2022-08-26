package main

import (
  "fmt"
  "os"
  "strings"
)

const (
  usage = `
string-man
  [command] [string]
  
  Available commands: lower, upper and title`
  err = "Unknown command: %q\n"
)

func main() {
  if len(os.Args) != 3 {
    fmt.Printf(usage)
    return
  }
  
  // main
  switch cmd, v := strings.ToLower(os.Args[1]), os.Args[2]; cmd {
  case "lower":
    v2 := strings.ToLower(v)
  case "upper":
    v2 := strings.ToUpper(v)
  case "title":
    v2 := strings.ToTitle(v)
  default:
    v2 := "Unknown command: " + v
  }
  
  fmt.Printf("%s\n", v2)
}
