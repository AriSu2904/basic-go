package main

import "fmt"

func main() {

	//Menambah data pada array secara langsung
	var numbers = [5]int{
		3, 4, 8, 2, 8,
	}

	//mengakses array pada index tertentu
	fmt.Println(numbers[1]) // 4
}
