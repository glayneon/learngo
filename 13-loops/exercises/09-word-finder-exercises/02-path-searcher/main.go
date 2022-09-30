package main

import (
    "fmt"
    "os"
    "strings"
)

const (
    usage = `
    This is gonna search the given path in PATH env variables.
    
    psearch [keyword] [keyword2] [keyword3] ...
    
    [keyword] : word that you want to find out in PATH env variables.
    
    `
    
    err1 = "There's no %q in PATH Env variable.\n"
)


func main() {
    pth := strings.Split(os.Getenv("PATH"), os.PathListSeparator)
    query := os.Args[1:]

    switch len(query) {
    case 0:
        fmt.Printf(usage)
        return
    }
    
queries:
    found := false
    for _, q := range query {
        q = strings.ToLower(q) // make q is lower-case
        
        for i, d := range pth {
            if q == d {
                fmt.Printf("#%-2d: %q\n", i+1, d)
                found = true
            }
        }
        if found != true {
            fmt.Printf(err1, q)
            continue queries
        }
    }

}
