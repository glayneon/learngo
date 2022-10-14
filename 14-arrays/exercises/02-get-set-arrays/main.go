package main

import "fmt"

func main() {
	var (
		names     [3]string  // The names of your three best friends
		distances [5]int     // The distances to five different locations
		data      [5]byte    // A data buffer with five bytes of capacity
		ratios    [1]float64 // Currency exchange ratios only for a single currency
		alives    [4]bool    // Up/Down status of four different web servers
		zero      [0]byte    // A byte array that doesn't occupy memory space
	)

    // assign to names
    names[0] = "Chase, Kim"
    names[1] = "Triss, You"
    names[2] = "Hans, Kim"
    
    // assign to distnaces
    distances[0] = 118
    distances[1] = 64
    distances[2] = 1133
    distances[3] = 443
    distances[4] = 17818
    
    // assign to data
    data[0] = "H"
    data[1] = "E"
    data[2] = "L"
    data[3] = "L"
    data[4] = "O"
    
    // assign to ratios
    ratios[0] = 1.41
    
    // assign to alives
    alives[0] = true
    alives[1] = true
    alives[2] = false
    alives[3] = true
    alives[4] = false
    
    // assign to zero
    zero[0] = "T"
    
    // print names array
    for i = 0; i < len(names)+1 ; i++ {
        fmt.Printf("%v", i)
    }
    
    
}
