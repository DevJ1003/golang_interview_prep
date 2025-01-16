// METHOD - 1
package main

import "fmt"

func reverseString(s string) string {
	n := len(s)
	result := ""

	// Iterate from the end of the string to the beginning
	for i := n - 1; i >= 0; i-- {
		result += string(s[i]) // Append each character to the result
	}

	return result
}

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	reversed := reverseString(input)
	fmt.Printf("Reversed string: %s\n", reversed)
}

//
//
//
//

// METHOD - 2
// package main

// import "fmt"

// func reverseString(s string) string {
//     reversed := ""
//     for _, char := range s {
//         reversed = string(char) + reversed
//     }
//     return reversed
// }

// func main() {
//     var input string
//     fmt.Print("Enter a string: ")
//     fmt.Scanln(&input)

//     reversed := reverseString(input)
//     fmt.Printf("Reversed string: %s\n", reversed)
// }
