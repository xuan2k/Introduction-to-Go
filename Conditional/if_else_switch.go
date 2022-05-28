package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 12
	if a < b {
		fmt.Println(a, "is smaller than", b)
	} else if a > b {
		fmt.Println(a, "is larger than", b)
	} else {
		fmt.Println(a, "is equal to", b)
	}

	option := 2
	switch option {
	case 1:
		fmt.Println("option 1 chosen!")
	case 2:
		fmt.Println("option 2 chosen!")
	case 3:
		fmt.Println("option 3 chosen!")
	default:
		fmt.Println("No case matched!!")
	}
}
