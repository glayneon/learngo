// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
  "fmt"
  "os"
)

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go i wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func main() {
  l := len(os.Args) - 1
  
  // conditional
  if l == 0 {
    fmt.Println("Give me Args")
  } else if l == 1 {
    fmt.Println("There is one: ", os.Args[1])
  } else if l == 2 {
    fmt.Println("There are two: ", os.Args[1], os.Args[2])
  } else {
    fmt.Printf("There are %v arguments", l)
  }
}
