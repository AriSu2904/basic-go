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

func getInfo(car CarInfo) {
	brand := car.GetBrand()
	engine := car.GetEngineType()
	fmt.Println("Brand :", brand)
	fmt.Println("Engine Type :", engine)
}

type Car struct {
	Name, Brand, Engine string
}

func (car Car) GetBrand() string {
	return car.Brand
}

func (car Car) GetEngineType() string {
	return car.Engine
}

type CarInfo interface {
	GetBrand() string
	GetEngineType() string
}
