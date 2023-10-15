package main

import "fmt"

func main() {

	//car := map[string]string{
	//	"type":  "Gasoline",
	//	"brand": "Toyota",
	//}
	var car = map[string]string{}
	car["type"] = "Gasoline"
	car["brand"] = "Toyota"

	fmt.Println(car)
	fmt.Println(car["type"])
	fmt.Println(car["brand"])

	var book = make(map[string]string)
	book["title"] = "Kosmos"
	book["author"] = "Carl Sagan"
	fmt.Println(book)
	fmt.Println(len(book))

	fmt.Println(book)
	fmt.Println(len(book))
}
