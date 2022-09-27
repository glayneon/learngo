package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	usage = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.
Wanna play?
`
	em1 = `
    wrong number. %q
`
	maxTurn = 10
)

func main() {
	verbose := false
	args := os.Args[1:]
	rand.Seed(time.Now().UnixNano())

	if len(args) < 1 {
		fmt.Printf(usage, maxTurn)
		return
	}

	if args[0] == "-v" {
		verbose = true
	}

	guess, err := strconv.Atoi(args[len(args)-1])
	if err != nil || guess < 0 {
		fmt.Printf(em1, args[1])
		return
	}

	for turn := 1; turn < maxTurn; turn++ {
		n := rand.Intn(guess + 1)
		if verbose == true {
			fmt.Printf("%2d ", n)
		}

		if n == guess {
			if turn == 1 {
				fmt.Println("Wow, In first turn, ")
			}
			fmt.Println("You Win!")
			return
		}
	}
	fmt.Println("You Lost... Try again?")
}
