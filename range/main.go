package main

import "fmt"

func main() {
	ids := []int{21, 22, 23, 24, 25}

	for i, id := range ids {
		fmt.Printf("%d - id: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("id: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	party := map[string]string{"Foo": "Fighter", "Bar": "Wizard"}
	for key, val := range party {
		fmt.Printf("%s: %s\n", key, val)
	}
}
