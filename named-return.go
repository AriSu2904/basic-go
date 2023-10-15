package main

import "fmt"

func main() {
	fmt.Println(name())
}

func name() (firstName string, lastName string) {
	firstName = "Ari"
	lastName = "Susanto"
	return
}
