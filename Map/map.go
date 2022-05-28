package main

import "fmt"

func main() {
	fruitShop := map[int]string{
		11: "Orange",
		12: "Apple",
		13: "Banana",
		14: "Durian",
	}
	fmt.Println(fruitShop)
	fmt.Println(fruitShop[13])
	fmt.Println(fruitShop[14])

	fmt.Println("Original Map:")
	fmt.Println(fruitShop)
	fruitShop[15] = "Grape"
	fmt.Println("Add more entry:")
	fmt.Println(fruitShop)
	fmt.Println("Remove the previous entry:")
	delete(fruitShop, 15)
	fmt.Println(fruitShop)

	for fruitID, fruitName := range fruitShop {
		fmt.Printf("Fruit ID: %d. Fruit Name: %s\n", fruitID, fruitName)

	}
}
