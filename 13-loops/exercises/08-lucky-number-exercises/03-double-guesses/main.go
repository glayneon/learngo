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
    wrong number. %q or %q
`
    em2 = `
    there would be a non-positive number... %q or %q
`
    maxTurn = 10
)

func main() {
    // get two command-line arguments
    args := os.Args[1:]
    rand.Seed(time.Now().UnixNano())
        
    if len(args) != 2 {
        fmt.Printf(usage, maxTurn)
        return
    }
    
    guess1, err1 := strconv.Atoi(args[0])
    guess2, err2 := strconv.Atoi(args[1])
    
    if err != nil || err2 != nil {
        fmt.Printf(em1, args[0], args[1])
        return
    } else if guess1 < 0 || guess2 < 0 {
        fmt.Printf(em2, args[0], args[1])
        return
    }
    rn := guess2
    for turn := 1; turn < maxTurn; turn++ {
        if guess1 > guess2 {
            rn = guess1
        }
        
        n := rand.Intn(rn + 1)
        
        if n == guess1 || n == guess2 {
            if turn == 1 {
                fmt.Println("Wow, It's first round!.")
            }
            fmt.Println("You Win!")
        }
    }
    fmt.Println("You Lost... Try again?")
}
