package main

import (
	"fmt"
)

// Recursive function to calculate factorial
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1) // Recursive call
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Factorial is not defined for negative numbers.")
	} else {
		fmt.Printf("The factorial of %d is %d\n", num, factorial(num))
	}
}

// Recursion is a programming technique where a function calls itself to solve
// a smaller instance of the same problem. This continues until the function
// reaches a base case, which is a condition that stops further recursion.
