package main

import (
	"fmt"
)

func main() {

	var num int
	fmt.Print("Enter the number of terms: ")
	fmt.Scan(&num)

	if num <= 0 {
		fmt.Println("Enter the possitive number!")
	}

	first, second := 0, 1
	fmt.Println("Fibonacci series: ")

	for i := 0; i <= num; i++ {
		fmt.Print(first, " ")
		next := first + second
		first = second
		second = next
	}
}
