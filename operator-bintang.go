package main

import "fmt"

func main() {

	dog := Animal{Name: "Anjing Kecil"}
	dog2 := &dog
	dog3 := &dog
	dog4 := &dog2

	*dog2 = Animal{Name: "Anjing besar"}

	fmt.Println(dog)
	fmt.Println(dog2)
	fmt.Println(dog3)
	fmt.Println(dog4)

}

type Animal struct {
	Name string
}
