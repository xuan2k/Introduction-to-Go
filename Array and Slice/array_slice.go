package main

import "fmt"

func main() {
	//Array
	var intArrayA = [3]int{1, 2, 3} //integer array with 3 elements: 1,2,3
	var intArrayB = [4]int{1}       //array with zero value
	var stringArray = [3]string{}   //array with zero value
	var boolArray = [2]bool{true}   //array with zero value
	fmt.Println(intArrayA)
	fmt.Println(intArrayB)
	fmt.Println(stringArray)
	fmt.Println(boolArray)

	var intArray = [5]int{1, 2, 3, 7, 8}
	for index, num := range intArray {
		fmt.Printf("%d th element: %d\n", index, num)
	}

	var floatArray = [5]float64{1, 2, 3, 7, 8}
	average := float64(0)
	for _, num := range floatArray {
		average += num
	}
	average /= float64(len(floatArray))
	fmt.Println("Average value of array is:", average)

	//Slice
	var intSlice = []int{1, 2, 3, 7, 8}
	fmt.Println("for loop:")
	for i := 0; i < len(intSlice); i++ {
		fmt.Printf("%d th element: %d\n", i, intSlice[i])
	}
	fmt.Println("for...range loop:")
	for index, num := range intSlice {
		fmt.Printf("%d th element: %d\n", index, num)
	}
}
