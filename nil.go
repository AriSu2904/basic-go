package main

import "fmt"

func main() {

	result := sayHello("Ari")
	if result != nil {
		fmt.Println(result)
	} else {
		fmt.Println("Hello, Nil!")
	}

}

func sayHello(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
