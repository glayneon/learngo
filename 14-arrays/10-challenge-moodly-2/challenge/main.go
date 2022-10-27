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
	"math/rand"
	"os"
	"time"
)

const (
	positive = 0
	negative = 1
)


func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("[your name] [positive|negative]")
		return
	}

	name, mood := args[0], args[1]

	moods := [...][3]string{
		{"happy 😀", "good 👍", "awesome 😎"},
		{"sad 😞", "bad 👎", "terrible 😩"},
	}

	// Generate Seed
	rand.Seed(time.Now().UnixNano())
	
	switch mood {
	case "positive":
		emotion := positive
	case "negative":
		emotion := negative
	default:
		fmt.Printf("There's no %q feeling in this.\n", mood)
		fmt.Println("[your name] [positive|negative]")
		return
	}
	
	n := rand.Intn(len(moods[emotion]))
	fmt.Printf("%s feels %s\n", name, moods[emotion][n])
}
