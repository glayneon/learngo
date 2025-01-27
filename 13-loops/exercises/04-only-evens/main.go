package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	even  = 2
	usage = `
    
    Usage : sumup [min] [max]
    
    [min] : It is an integer and it would be shorter than [max]
    [max] : It is an integer and it would be greater than [min]
    
`
	errm1 = "%q is not an integer.\n"
	errm2 = "Missing Arguments.\n"
	errm3 = "[min] %q must be shorther than [max].\n"
)

func main() {
	args := os.Args[1:]

	// check command-line parameters
	if l := len(args); l != 2 {
		fmt.Printf(errm2)
		fmt.Printf(usage)
		return
	}

	// get the min and max values and try to convert to an integer
	min, err1 := strconv.Atoi(args[0])
	max, err2 := strconv.Atoi(args[1])

	// check there would be an error occured.
	if err1 != nil {
		fmt.Printf(errm1, args[0])
		fmt.Printf(usage)
		return
	} else if err2 != nil {
		fmt.Printf(errm1, args[1])
		fmt.Printf(usage)
		return
	} else if min >= max {
		fmt.Printf(errm3, args[0])
		fmt.Printf(usage)
		return
	}

	var sum int
	for i := min; i <= max; i++ {
		// fmt.Print(i)

		if i%even == 0 {
			fmt.Print(i)
			if i < max-1 {
				fmt.Print(" + ")
			}
			sum += i
		}
	}
	fmt.Printf(" = %d\n", sum)
}
