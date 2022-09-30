package main

import (
    "fmt"
    "os"
    "strings"
)

const corpus = "lazy cat jumps again and again and again"

func main() {
    word := strings.Fields(corpus)
    args := os.Args[1:]
    
// label
queries:
    for _, q range args {
        
    search:
        for i, v range word {
            switch v {
            case "and", "or", "the":
                break search
            }
            
            if q == v {
                fmt.Printf("#%-2d: %q\n", i+1, v)
                continue queries
            }
        }
    }
}
