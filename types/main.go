package main

import (
	"fmt"
	"strconv"
)

var name string = "Foo"
var age int = 30

func main() {

	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// unint uint8 uint16 uint32 uint64
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	size, email := 1.3, "foo@gmail.com" //sharthand for local variables

	// Type conversion
	// https://golang.org/pkg/strconv/
	// strconv.Atoi (string to int)
	// strconv.Itoa (int to string)
	age := strconv.Itoa(age) + "y"

	// fmt.Printf("%T\n", name) // get type
	fmt.Println(name, age, size, email)
}
