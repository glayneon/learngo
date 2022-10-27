package main

import (
    "fmt"
    "strings"
)


func main() {
    var (
        upper [3]string
        lower [3]string
    )
    
    books := [3]string{
        "Kafka's Revenge",
        "Stay Golden",
        "Everythingship",
    }
    
    sep := "\n" + strings.Repeat("~", 30) + "\n"
    
    // initiate upper
    for i, v := range books {
        upper[i] = strings.Upper(books[i])
    }
    
    // initiate lower
    for i, v := range books {
        lower[i] = strings.Lower(books[i])
    }
    
    // print them all
    fmt.Printf("books: %q\n", books)
    fmt.Printf("upper: %q\n", upper)
    fmt.Printf("lower: %q\n", lower)
}
