package main

import "fmt"

func main() {

	//Declaration Variable
	var name string
	name = "Ari"

	//Using alias
	var myAge byte // alias for int64 depend on OS Bit
	myAge = 18

	// Initialization Variable
	var name1 = "Ari"

	//Short Declaration
	name2 := "Ari"

	//Multiple Declaration
	var x, y, z int32
	//or
	var (
		firstName = "Ari"
		lastName  = "Susanto"
	)

	var (
		asd string
		dsa string
	)

	asd = "asd"
	dsa = "asd"

	fmt.Println("Hello,", name, name1, name2, myAge, x, y, z, firstName, lastName, asd, dsa)
}
