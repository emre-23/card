package main

import (
	"fmt"
)

func main() {
	var slices = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, slice := range slices {
		if i%2 == 0 {
			fmt.Println(slice, "is even")
		} else {
			fmt.Println(slice, "is odd")
		}

	}
}
