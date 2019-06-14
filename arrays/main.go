package main

import "fmt"

func main() {

	items := [3]string{"Torch", "Key", "Sword"}

	fmt.Println(items)
	fmt.Println(items[0])
	fmt.Println(len(items)) // length
	fmt.Println(items[1:3]) // range
}
