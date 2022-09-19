package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    const fmt = "\r%s "
    const dur = 250
    const msg = "Please Wait. Processing....\n"
    
    for {
        rand.Seed(time.Now().UnixNano())
        
        ln := rand.Intn(4)
        switch ln {
        case 0:
            rc := "/"
        case 1:
            rc := "|"
        case 2:
            rc := "\"
        case 3:
            rc := "-"
        }
        // print
        fmt.Printf(fmt, pc)
        fmt.Printf(msg)
        
        time.Sleep(dur * time.Millisecond)
    }
}
