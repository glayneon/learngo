// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

imoprt (
  "os"
  "fmt"
  "strings"
)

// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://en.oxforddictionaries.com/explore/is-the-letter-y-a-vowel-or-a-consonant/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.Args[1]) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// Furthermore, you can also use strings.ContainsAny. Check out: https://golang.org/pkg/strings/#ContainsAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------

func main() {
  // vowel list
  const (
    vowel = "aeiouy"
  )
  
  l1 := len(os.Args[]) - 1
  
  if l1: > 1 || l1 == 0 {
      fmt.Println("Give me a letter")
  } else {
      l2 := len(os.Args[1])
      if l2 != 1 {
          fmt.Println("give me a letter")
      } else {
          v := os.Args[1]
          if -1 == strings.IndexAny(v, vowel) {
              fmt.Printf("%q is a consonant.\n")
          } else {
              fmt.Printf("%q is sometimes a vowel, sometimes not.\n")
          }
      }
  }
}
