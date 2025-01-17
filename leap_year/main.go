package main

import (
	"fmt"
)

// nested if statement
// func main() {

// 	var year int

// 	fmt.Print("Enter the year to check for leap year: ")
// 	fmt.Scan(&year)

// 	if year%100 == 0 {
// 		if year%400 == 0 {
// 			fmt.Printf("Year %d is a leap year!", year)
// 		} else {
// 			fmt.Printf("Year %d is not a leap year!", year)
// 		}
// 	} else {
// 		if year%4 == 0 {
// 			fmt.Printf("Year %d is a leap year!", year)
// 		} else {
// 			fmt.Printf("Year %d is not a leap year!", year)
// 		}
// 	}
// }

// boolean if-else
func main() {

	var year int

	fmt.Print("Enter the year to check for leap year: ")
	fmt.Scan(&year)

	if year%100 == 0 && year%400 == 0 || year%100 != 0 && year%4 == 0 {
		fmt.Printf("Year %d is a leap year!", year)
	} else {
		fmt.Printf("Year %d is not a leap year!", year)
	}
}
