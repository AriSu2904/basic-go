package main

import "fmt"

func main() {

	var numbers = [5]int{
		3, 4, 8, 2, 8,
	}

	for i, number := range numbers {
		fmt.Println(i, number)
	}

	cars := map[string]string{
		"type":  "Gasoline",
		"brand": "Toyota",
	}

	for i, car := range cars {
		fmt.Println(i, car)
	}

	for _, car := range cars {
		fmt.Println(car)
	}
}
