package main

import "fmt"

func main() {
	result := firstName()

	switch result.(type) {
	case string:
		fmt.Println("Return adalah string")
		break
	case int:
		fmt.Println("Return adalah integer")
	case bool:
		fmt.Println("Return adalah boolean")
	}

}

func firstName() interface{} {
	return true
}
