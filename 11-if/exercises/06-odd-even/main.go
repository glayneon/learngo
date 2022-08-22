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
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Odd or Even
//
//  1. Get a number from the command-line.
//
//  2. Find whether the number is odd, even and divisible by 8.
//
// RESTRICTION
//  Handle the error. If the number is not a valid number,
//  or it's not provided, let the user know.
//
// EXPECTED OUTPUT
//  go run main.go 16
//    16 is an even number and it's divisible by 8
//
//  go run main.go 4
//    4 is an even number
//
//  go run main.go 3
//    3 is an odd number
//
//  go run main.go
//    Pick a number
//
//  go run main.go ABC
//    "ABC" is not a number
// ---------------------------------------------------------

const (
	divider = 8
	pr1     = "%d is an even number and it's divisible by %d\n"
	pr2     = "%d is an even number\n"
	pr3     = "%d is an odd number\n"
	err1    = "Pick a number\n"
	err2    = "%q is not a number\n"
)

func main() {
	// short statement
	if a := os.Args; len(a) != 2 {
		fmt.Printf(err1)
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf(err2, a[1])
	} else {
		if n%2 == 0 && n%8 == 0 {
			fmt.Printf(pr1, n, divider)
		} else if n%2 == 0 {
			fmt.Printf(pr2, n)
		} else {
			fmt.Printf(pr3, n)
		}
	}
}
