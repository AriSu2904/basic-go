package main

import "fmt"

func main() {
	type idCard string

	var myId idCard = "12312313"

	fmt.Println(myId)
}
