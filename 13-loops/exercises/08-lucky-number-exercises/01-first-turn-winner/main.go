package main

import (
    "fmt"
    "time"
    "math/rand"
    "os"
    "strconv"
)

const (
    usgage = `
    randguess [number]
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
        fmt.Printf(usage)
        return
    }
    
    guess, err := strconv.Atoi(args[0])
    if err != nil || guess < 0 {
        fmt.Printf(em1, args[0])
        return
    }
    
    for turn := 0; turn < maxTurn; turn++ {
        n := rand.Intn(guess + 1)
        
        switch {
        case turn == 0:
            fmt.Println("Wow, In first turn,")
            fallthrough
        case n == guess:
            fmt.Println("You Win!")
            return
        }
    }
    fmt.Println("You Lost... Try again?")
}
