package main

import "fmt"

type Car struct {
	Brand string
	Color string
}

func main() {
	type customerID string
	type fuelCapacity float64
	var cus1 customerID = "12345"
	var car1Cap fuelCapacity = 40
	fmt.Println("customer's ID:", cus1)
	fmt.Println("fuel capacity of car1:", car1Cap)

	car1 := Car{Brand: "Suzuki", Color: "blue"}
	fmt.Println(car1)

	car2 := Car{Brand: "Mercedes", Color: "black"}
	car2.Brand = "Toyota"
	car2.Color = "white"
	fmt.Println(car2)
}
