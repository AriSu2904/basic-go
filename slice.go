package main

import "fmt"

func main() {

	//var num = make([]int, 2, 5)
	//num[0] = 1
	//num[1] = 2
	//
	//
	//var slicedNumber1 = numbers[4:8]
	//
	//var slicedNumber2 = numbers[5:8]

	//fmt.Println("Sebelum dirubah")
	//fmt.Println(slicedNumber1)
	//fmt.Println(slicedNumber2)
	//slicedNumber1[2] = 100
	//fmt.Println("Setelah dirubah")
	//fmt.Println(slicedNumber1)
	//fmt.Println(slicedNumber2)
	//fmt.Println(numbers)

	var numbers = [5]int{1, 2, 3, 4, 5}

	var slicedNum = numbers[1:4]
	var slicedNum2 = append(slicedNum, 100)
	fmt.Println(slicedNum2)
	fmt.Println(slicedNum)
	fmt.Println(numbers)

}
