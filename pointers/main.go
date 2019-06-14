package main

import "fmt"

func main() {
	//Pointers: passing reference to address (or pointer) instead of the value of the data itself improves performance

	a := 5
	b := &a //0xc000080008 memory address

	fmt.Println(a, b)
	fmt.Printf("%T\n", b) //*pointer

	// Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val via pointer
	*b = 10
	fmt.Println(b, *b, a)
}
