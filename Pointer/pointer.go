package main

import "fmt"

func main() {
	var i int = 5
	var pI *int = &i
	fmt.Println("Value of i is", i)
	fmt.Println("Address of i is", &i)
	fmt.Println("Address of i is", pI)
	fmt.Printf("Type of i: %T\nType of pI: %T\n", i, pI)
}
