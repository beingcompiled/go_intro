package main

import "fmt"

func main() {
	// Define map
	// party := make(map[string]string)
	// party["Foo"] = "Fighter"
	// party["Bar"] = "Wizard"

	party := map[string]string{"Foo": "Fighter", "Bar": "Wizard"}

	fmt.Println(party)
	fmt.Println(len(party))
	fmt.Println(party["Foo"])

	// Delete
	delete(party, "Foo")
	fmt.Println(party)
}
