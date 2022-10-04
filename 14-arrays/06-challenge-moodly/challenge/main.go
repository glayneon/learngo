package main

import (
    "fmt"
    "os"
    "time"
    "math/rand"
)

const (
    usage = `
    [your name]
`
    fmt1 = "%s feels %s\n"
)

func main() {
    var feeling string
    
    react := [...]string{
        "good",
        "bad",
        "sad",
        "happy",
        "awesome",
        "terrible",
    }
    
    rand.Seed(time.Now().UnixNano())
    
    if len(os.Args) != 2 {
        fmt.Println(usage)
        return
    }
    
    feeling = react[rand.Intn(len(react))]
    // print feeling
    fmt.Printf(fmt1, os.Args[1], feeling)
}

    
