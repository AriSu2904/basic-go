package main

import "fmt"

func main() {
	access(0)
}

func access(pin int) {
	defer logout()
	fmt.Println("Welcome")
	fmt.Println(10 / pin)
}

func logout() {
	fmt.Println("Logout...")
}
