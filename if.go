package main

import "fmt"

func main() {

	//name := "Arsi"
	//
	//if name == "Ari" {
	//	fmt.Println("Hi", name)
	//} else {
	//	fmt.Println("Hi, There...")
	//}

	age := 18

	if age == 17 {
		fmt.Println("Your 17 y.o")
	} else if age == 18 {
		fmt.Println("Your 18 y.o")
	} else {
		fmt.Println("How old are u?")
	}

	if name := "ari"; name == "ari" {
		fmt.Println("Hello", name)
	}
}
