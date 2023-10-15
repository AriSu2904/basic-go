package main

import "fmt"

func main() {
	runApp(false)
}

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Logout....")
}

func runApp(admin bool) {
	defer endApp()
	if !admin {
		panic("You are not admin!")
	}
}
