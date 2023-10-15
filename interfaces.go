package main

import "fmt"

func main() {

	toyota := Car{
		Name:   "Supra",
		Brand:  "Toyota",
		Engine: "Electric",
	}

	getInfo(toyota)
}

type CarInfo interface {
	GetBrand() string
	GetEngineType() string
}

func getInfo(car CarInfo) {
	fmt.Println(car.GetBrand())
	fmt.Println(car.GetEngineType())
}

func (car Car) GetBrand() string {
	return car.Brand
}

func (car Car) GetEngineType() string {
	return car.Engine
}

type Car struct {
	Name, Brand, Engine string
}
