package main

import "fmt"

func reverseNumber(num int) int {
	reversed := 0
	for num != 0 {
		digit := num % 10              // Extract the last digit
		reversed = reversed*10 + digit // Append it to the reversed number
		num /= 10                      // Remove the last digit
	}
	return reversed
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	reversed := reverseNumber(num)
	fmt.Printf("Reversed number: %d\n", reversed)
}
