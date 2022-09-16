package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    const fmt = "\r%s "
    const lc = "\/|-"
    const dur = 250
    const msg = "Please Wait. Processing....\n"
    
    for {
        rand.Seed(time.Now().UnixNano())
        pc := lc[rand.Intn(len(lc)+1)]
        
        // print
        fmt.Printf(fmt, pc)
        fmt.Printf(msg)
        
        time.Sleep(dur * time.Millisecond)
    }
}
