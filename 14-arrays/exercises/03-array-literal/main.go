package main

import "fmt"

func main() {
	var (
        names     [3]string = [3]string{
            "Chase, Kim",
            "Triss, You",
            "Hans, Kim",
        }
        distances [5]int = [5]int{118, 64, 1133, 443, 17818}
        data      [5]byte = [5]byte{"H", "E", "L", "L", "O"}
        ratios    [1]float64 = [1]float64{1.14}
        alives    [4]bool = 4[bool]{
            true,
            true,
            false,
            false,
            true,
        }
		zero      [0]byte    // A byte array that doesn't occupy memory space
	)

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
    for i = 0; i < len(names) ; i++ {
		fmt.Printf("names[%v]: %v\n", i, names[i])
    }
	
	for i = 0; i < len(distances); i++ {
		fmt.Printf("distance[%v]: %v\n", i, distances[i])
	}
	
	for i = 0; i < len(data); i++ {
		fmt.Printf("data[%v]: %v\n", i, data[i])
	}
	
	for i = 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%v]: %v\n", i, ratios[i])
	}
	
	for i = 0; i < len(alives); i++ {
		fmt.Printf("alives[%v]: %v\n", i, alives[i])
	}
	
	// for range
	for i, v := range(names) {
		fmt.Printf("names[%v]: %v\n", i, v)
	}
	
	for i, v := range(distances) {
		fmt.Printf("distances[%v]: %v\n", i, v)
	}
}
