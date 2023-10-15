package main

import "fmt"

func main() {

	numbers := 10

	for i := 1; i <= numbers; i++ {
		fmt.Print(i, " ")
		if i == 5 {
			break
		}
	}
}
