package main

import (
    "fmt"
    "time"
    "math/rand"
    "os"
    "strconv"
)

const (
    usgage = `Welcome to the Lucky Number Game! 🍀

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
    args := os.Args[1:]
    rand.Seed(time.Now().UnixNano())
        
    if len(args) != 2 {
        fmt.Printf(usage, maxTurn)
        return
    }
    
    guess, err := strconv.Atoi(args[1])
    if err != nil || guess < 0 {
        fmt.Printf(em1, args[1])
        return
    }
    
    if args[0] != "-v" 
    
    for turn := 1; turn < maxTurn; turn++ {
        n := rand.Intn(guess + 1)
        
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
