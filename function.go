package main

import "fmt"

func main() {
	myFunction("Ari")    //Hello My Name Is Ari
	myFunction("Yunita") //Hello My Name Is Yunita
	fullName("Ari", "Susanto")
}

func myFunction(name string) {
	fmt.Println("Hello My Name Is", name)
}

func fullName(firstName string, lastName string) {
	full := firstName + " " + lastName
	fmt.Println(full)
}
