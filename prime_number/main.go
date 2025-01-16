package main

import (
	"fmt"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false // Numbers less than or equal to 1 are not prime
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false // Divisible by a number other than 1 and itself
		}
	}
	return true // Prime number
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if isPrime(num) {
		fmt.Printf("%d is a prime number.\n", num)
	} else {
		fmt.Printf("%d is not a prime number.\n", num)
	}
}
