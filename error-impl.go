package main

import (
	"errors"
	"fmt"
)

func main() {

	res, err := charge(-100)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Transaction Status :", res)
	}
}

func charge(balance int) (bool, error) {
	if balance <= 0 {
		return false, errors.New("Error!\nMinimal Topup 1 Rupiah!")
	} else {
		return true, nil
	}
}
