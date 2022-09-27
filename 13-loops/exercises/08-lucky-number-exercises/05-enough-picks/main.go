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
    maxNum = 10
)

func main() {
    args := os.Args[1:]
    rand.Seed(time.Now().UnixNano())
        
    if len(args) < 1 {
        fmt.Printf(usage, maxTurn)
        return
    }
    
    guess, err := strconv.Atoi(args[0])
    if err != nil || guess < 0 {
        fmt.Printf(em1, args[0])
        return
    }
    
    maxNum := 10
    if guess > maxNum {
        maxNum = guess
    }
        
    for turn := 1; turn < maxTurn; turn++ {
        n := rand.Intn(maxNum + 1)
        fmt.Printf("3%d" n)
        
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
