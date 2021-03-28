package main

import "fmt"

func BracketCombinations(num int) int {

	// code goes here
	if num < 0 {
		return 0
	} else if num <= 1 {
		return 1
	}

	arr := make([]int, num+1)
	arr[0] = 1
	arr[1] = 1

	for i := 2; i <= num; i++ {
		ins := i - 1
		out := 0
		for ins >= 0 {
			arr[i] += arr[ins] * arr[out]
			ins--
			out++
		}
	}

	return arr[num]

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(BracketCombinations(readline()))

}
