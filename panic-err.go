package main

import "fmt"

func main() {
	accessible(false)
}

func accessible(admin bool) {
	defer signOut()
	fmt.Println("Welcome...")
	if !admin {
		panic("Your not allowed here!")
	}
}

func signOut() {
	fmt.Println("Logout...")
}
