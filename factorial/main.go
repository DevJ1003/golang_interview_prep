package main

import (
	"fmt"
)

func main() {

	var num int
	var factorial int = 1
	fmt.Print("Enter the number: ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Factorial not defined for negative numbers.")
		return
	}

	for i := num; i > 0; i-- {
		factorial = factorial * i
	}

	fmt.Printf("The factorial of %d is %d", num, factorial)
}
