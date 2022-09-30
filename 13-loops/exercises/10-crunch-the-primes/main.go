package main

import (
    "fmt"
    "os"
    "strconv"
)

var primcount int = 0

func primeFunc(primNum int) int {
    for i := 2; i< primNum/2, i++ {
        if primeNum%i == 0 {
            primcount++
        }
    }
    return primecount
}

func main() {
    nums := os.Args[1:]
  
    search:
    for _, n := range nums {
        if n, err := strconv.Atoi(n); err != nil {
            continue search
        }
        
        // check prime number
        primcount = primeFunc(n)
        
        if primcount == 0 && n != 1 {
            fmt.Printf("%-3d", n)
        }
    }
}
        
