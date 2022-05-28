package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		fmt.Println("i is:", i)
	}
	for i := 1; i <= 4; i++ {
		fmt.Println("i is:", i)
		fmt.Println("this line is not skipped!")
		continue
		// Unreachable line
		//fmt.Println("this line is skipped!!")
	}
	for i := 1; i <= 4; i++ {
		fmt.Println("i is:", i)
		if i == 2 {
			fmt.Println("Loop skipped!! Jump out of loop !!")
			break
		}
	}
}
