package main

import "fmt"

func Hello() {
	fmt.Println("Hello World")
}

func triangleArea(height float64, edge float64) {
	area := (height * edge) / 2.0
	fmt.Printf("Area of the triangle with height %.2f and edge %.2f is: %.2f\n", height, edge, area)
}

func rectangleArea(height float64, width float64) float64 {
	area := height * width
	return area
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	Hello()
	triangleArea(3.0, 5.5)

	recArea := rectangleArea(9.0, 12.3)
	fmt.Printf("Area of the rectangle is: %.2f\n", recArea)

	a := 10
	b := 5
	fmt.Printf("Origin value of a = %d and b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("Exchanged value a = %d and b = %d\n", a, b)
}
