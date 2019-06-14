package main

import "fmt"

func main() {

	x := 10
	y := 10

	// If else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is more than %d\n", y, x)
	}

	// Else if
	color := "blue"

	if color == "red" {
		fmt.Printf("color is red\n")
	} else if color == "blue" {
		fmt.Printf("color is blue\n")
	} else {
		fmt.Printf("color is NOT blue or red\n")
	}

	// Switch
	switch color {
	case "red":
		fmt.Printf("color is red\n")
	case "blue":
		fmt.Printf("color is blue\n")
	}

	// && || operators
}
