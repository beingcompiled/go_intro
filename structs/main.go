package main

import "fmt"

// define struct
type Person struct {
	// name      string
	// class     string
	// race      string
	name, class, race string
	hitpoints         int
}

// Value Receiver Method
func (p Person) greet() string {
	return "Greetings Traveller, I am " + p.name + " the " + p.class + " " + p.race + "."
}

// Pointer Receiver Method
func (p *Person) isDamaged(val int) {
	p.hitpoints -= val
}

func main() {
	// person1 := Person{name: "Foo", race: "Elf", class: "Fighter", hitpoints: 20}
	person1 := Person{"Foo", "Fighter", "Elf", 20}

	fmt.Println(person1)
	fmt.Println(person1.name)
	person1.isDamaged(10)
	person1.isDamaged(3)
	fmt.Println(person1.hitpoints)
	fmt.Println(person1.greet())
}
