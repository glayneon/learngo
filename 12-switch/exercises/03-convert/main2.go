package main

import (
	"fmt"
	"os"
)
// ---------------------------------------------------------
// EXERCISE: Convert
//
//  Convert the if statement to a switch statement.
// ---------------------------------------------------------

const (
	usage       = "Usage: [username] [password]"
	errUser     = "Access denied for %q.\n"
	errPwd      = "Invalid password for %q.\n"
	accessOK    = "Access granted to %q.\n"
	user, user2 = "jack", "inanc"
	pass, pass2 = "1888", "1879"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	//
	// REFACTOR THIS TO A SWITCH
	//
	switch {
	case u == user && p == pass:
		fallthrough
	case u == user2 && p == pass2:
		fmt.Printf(accessOK, u)
	case u == user && p != pass:
		fallthrough
	case u == user2 && p != pass2:
		fmt.Printf(errPwd, u)
	default:
		fmt.Printf(errUser, u)
	}
}
