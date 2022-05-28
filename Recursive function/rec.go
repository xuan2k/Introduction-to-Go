package main

import "fmt"

//recursive version calculate sum of a slice
func recurSumOfSlice(numSlice []float64, index int) float64 {
	if index == len(numSlice)-1 {
		return numSlice[index]
	}
	return numSlice[index] + recurSumOfSlice(numSlice, index+1)
}

func main() {
	nums := []float64{1.5, 2.7, 3, 5, 6.5}
	fmt.Println("Sum of slice: ", recurSumOfSlice(nums, 0))
}
