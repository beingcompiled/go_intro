package main

import ("fmt"
		"math/rand")

func main() {
	fmt.Println("Hello random", rand.Intn(100) ) // capitalized functions are exported
}