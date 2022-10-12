package main

import (
    "fmt"
)


func main() {

    names := [...]string{
        "Luke Kim",
        "Triss You",
        "Eathan Tim",
    }
    
    distances := [...]int32{
        1153,
        830,
        251,
        3045,
        3557,
    }
    
    var (
        data [5]byte
    )
    
    ratios := [2]float32{
        1,
        +0.41,
    }
    alives := [3]string{
        "up",
        "up",
        "down",
    }
    
    var zero [0]byte
    
    for _, v := range(names) {
        fmt.Printf("%q")
    }
    
    
        
}
