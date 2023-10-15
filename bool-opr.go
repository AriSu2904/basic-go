package main

import "fmt"

func main() {
	var point = 80
	var rating = 9

	var passPoint = point >= 70
	var passRating = rating >= 8

	var isPass = passPoint && passRating

	fmt.Println(isPass) //true
}
