package main

import (
	"fmt"
	"introduction/config"
)

func main() {
	result := config.GetInfo()
	fmt.Println(result)
}
