package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// taking input from the user
	n := os.Args[1]
	// converting string into integer
	N, _ := strconv.Atoi(n)
	fmt.Println(N)

	// using for loop for solving the problem time taken 3.445s
	for N > 1 {
		if N%2 == 0 {
			N = N / 2
			fmt.Println(N)
		} else {
			N = N*3 + 1
			fmt.Println(N)
		}
	}

	// using switch case for solving the problem time taken 3.320s
	switch N > 1 {
	case N%2 == 0:
		N = N / 2
		fmt.Println(N)
	case N%2 == 1:
		N = N*3 + 1
		fmt.Println(N)
	default:
		fmt.Println("Enter the correct positive number")
	}
}
