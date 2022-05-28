package main

import "fmt"

type Name string
type workHours int
type Employee struct {
	name Name
	wh   workHours
}

func (n Name) greeting(name1 Name) {
	fmt.Println("Hello", name1, "!")
}

func (h workHours) getSalary(wh workHours) {
	salary := wh * 20
	fmt.Println("Your salary is:", salary)
}

func main() {
	var e1 Employee
	e1.name = "Xuan"
	e1.wh = 50
	e1.name.greeting(e1.name)
	e1.wh.getSalary(e1.wh)
}
