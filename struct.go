package main

import "fmt"

func main() {
	//var john Person
	//john.FirstName = "John"
	//john.LastName = "Doe"
	//john.Address = "California"
	//john.Age = 45
	//
	//fmt.Println(john.FirstName)

	john := Person{
		FirstName: "John",
		LastName:  "Doe",
		Address:   "California",
		Age:       45,
	}
	john.waving("Ari")
}

func (person Person) waving(name string) {
	fmt.Println("Hello", name, "My name is", person.FirstName)
}

type Person struct {
	FirstName, LastName, Address string
	Age                          int
}
