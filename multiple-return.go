package main

import "fmt"

func main() {
	name, age := info("Ari", 18)
	//kita bisa gunakan _ untuk mengabaikan data tertentu
	nama, _ := info("Ari", 18)

	fmt.Println(name, age)
	fmt.Println(nama)

}
func info(name string, age int) (string, int) {
	return name, age
}
