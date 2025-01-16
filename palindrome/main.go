package main

import "fmt"

func isPalindrome(s string) bool {
	reversed := ""

	// Reverse the string
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}

	// Check if the original string is the same as the reversed string
	return s == reversed
}

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println(input, "is a palindrome.")
	} else {
		fmt.Println(input, "is not a palindrome.")
	}
}
