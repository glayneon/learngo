package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

const (
    fmt1 = "%5s"
    fmt2 = "%5d"
    usage = `
    
    Usage : metrix [ %q ] [size]
    
    [size] : it is greater than zero.
    
`
    validOps = "%*/+-"
    invalidOps = -1
    errm1 = `
    
    Size is missing.
`
    errm2 = `
    
    Invalid Operation.
`
)


func main() {
    args := os.Args[1:]

    switch l := len(args) {
    case l == 1: {
        fmt.Printf(errm1)
        fallthrough
        }
    case l > 1: {
        fmt.Printf(usage, validOps)
        return
        }
    }
    
    op := args[0]
    if strings.IndexAny(op, validOps) == invalidOps {
        fmt.Printf(errm2)
        fmt.Printf(usage, vaildOps)
        return
    }
    
    size, err := strconv.Atoi(args[1])
    if err != nil || size <= 0 {
        fmt.Printf(errm1)
        fmt.Printf(usage, validOps)
        return
    }
    
    // Print Header
    fmt.Printf(fmt1, op)
    for i := 0; i <= size; i++ {
        fmt.Printf(fmt2, i)
    }
    
    fmt.Println()
    
    // Print metrix
    for i := 0; i <= size; i++ {
    }
            
}
