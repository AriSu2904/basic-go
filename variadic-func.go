package main

import "fmt"

func main() {
	varargsExample(10, 10, 10)
}

func varargsExample(numbers ...int) {
	var temp int
	for _, number := range numbers {
		temp += number
	}

	fmt.Println(temp)
}
