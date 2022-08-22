package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go 2024
//    2024 is a leap year.
// ---------------------------------------------------------

// constant
const (
	err1 = "Give me a year number\n"
	err2 = "%q is not a valid year.\n"
	pr1  = "%d is not a leap year.\n"
	pr2  = "%d is a leap year.\n"
)

func main() {
	if a := os.Args; len(a) != 2 {
		fmt.Printf(err1)
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf(err2, a[1])
	} else {
		if n%4 == 0 {
			fmt.Printf(pr2, n)
		} else {
			fmt.Printf(pr1, n)
		}
	}
}
