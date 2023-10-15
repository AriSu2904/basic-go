package main

import "fmt"

func main() {
	fmt.Println(example(1)) // Output : true
}

func example(num int) interface{} {
	if num == 1 {
		return true
	} else if num == 2 {
		return "Dua"
	} else {
		return 0
	}
}
