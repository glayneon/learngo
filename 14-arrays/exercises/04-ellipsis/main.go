package main

import "fmt"

func main() {
  names := [...]string{
    "Chase, Kim",
    "Triss, You",
    "Hans, Kim",
  }
  distances := [...]int{118, 64, 1133, 443, 17818}
  data := [5]byte{"H", "E", "L", "L", "O"}
  ratios := [1]float64{1.14}
  alives := [5]bool{
            true,
            true,
            false,
            false,
            true,
  }
    
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
