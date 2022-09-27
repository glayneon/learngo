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
        
    if len(args) != 1 {
        fmt.Printf(usage, maxTurn)
        return
    }
    
    guess, err := strconv.Atoi(args[0])
    if err != nil || guess < 0 {
        fmt.Printf(em1, args[0])
        return
    }
    
    for turn := 1; turn < maxTurn; turn++ {
        n := rand.Intn(guess + 1)
        
        switch {
        case turn == 0:
            fmt.Println("Wow, In first turn,")
            fallthrough
        case n == guess:
            rand.Seed(time.Now().UnixNano())
            switch rand.Intn(3); {
            case 0:
                fmt.Println("Winner!!")
            case 1:
                fmt.Println("You Win!")
            case 2:
                fmt.Println("Awesome!")
            }
            return
        }
    }
    rand.Seed(time.Now().UnixNana())
    switch rand.Intn(3); {
    case 0:
        fmt.Println("Loser!!")
    case 1:
        fmt.Println("You lost..")
    case 2:
        fmt.Println("Seriously??")
    }
}
