package main

import "fmt"

func main() {

	point := 80

	switch {
	case point > 80:
		fmt.Println("Great!")
		break
	case point > 70:
		fmt.Println("Good Job!")
	case point > 60:
		fmt.Println("Nice Try!")
	default:
		fmt.Println("......")
	}

}
