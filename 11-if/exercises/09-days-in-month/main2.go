package main

import (
    "os"
    "fmt"
    "time"
    "strings"
)


func main() {
    // check the commandline parameters
    if a := os.Args; len(a) != 2 {
        fmt.Printf("Give me a month. \n")
        return
    }

    // get this year and check whether leap year or not
    year := time.Now().Year()
    leap := year%4 == 0 && ( year%100 != 0 || year%400 != 0)
    month := strings.ToLower(os.Args[1])
    
    // check the month
    if m := month; m == "april" ||
            m == "june" ||
            m == "september" ||
            m == "november" {
            day = 30
    } else if m == "january" || 
            m == "march" ||
            m == "may" ||
            m == "july" ||
            m == "august" ||
            m == "october" ||
            m == "decem ber" {
            day = 31
    } else if m == "febrary" {
        if leap {
            day = 29
        } else {
            day = 28
        }
    } else {
        fmt.Printf("%q is not a month.\n")
        return
    }

    fmt.Printf("%q has %d days.\n", month, day)
}
    
