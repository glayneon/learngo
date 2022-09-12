package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	usage = `
  Give me the size of the table

`

	merr = "Wrong size\n"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf(usage)
		return
	}

	// check it could be converted to integer
	if t, err := strconv.Atoi(os.Args[1]); err != nil || t < 0 {
		fmt.Printf(merr)
	} else {
		// Print Headers
		fmt.Printf("%5s", "X")

		// Print Headers
		for i := 0; i <= t; i++ {
			fmt.Printf("%5d", i)
		}
		fmt.Println()

		// Print contents of the table
		for i := 0; i <= t; i++ {
			fmt.Printf("%5d", i)
			for j := 0; j <= t; j++ {
				fmt.Printf("%5d", i*j)
			}
			// new line
			fmt.Println()
		}
	}
}
