package main

import (
	"fmt"
)

func main() {
	
	plantCapacities := []float64{30, 30, 30, 60, 60, 100}
	
	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Print("Please choose an option: ")
	
	var option string
	
	fmt.Scanln(&option)
	
	switch option {
		case "1":
			for idx, cap := range plantCapacities {
				fmt.Printf("Plant %d capacity: %.0f\n", idx, cap)
			}
	}
}