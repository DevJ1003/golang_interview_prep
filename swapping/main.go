package main

import (
	"fmt"
)

// without third variable
// func main() {
// 	var x, y int

// 	fmt.Print("Enter value for x & y: ")
// 	fmt.Scan(&x, &y)
// 	fmt.Printf("\nValues before swapping: %d %d\n", x, y)

// 	x = x + y
// 	y = x - y
// 	x = x - y

// 	fmt.Printf("Values after  swapping: %d %d\n", x, y)
// }

// with third variable
func main() {
	var x, y, z int
	fmt.Print("Enter the values for x & y: ")
	fmt.Scan(&x, &y)
	fmt.Printf("\nValues before swapping: %d %d\n", x, y)

	z = x
	x = y
	y = z

	fmt.Printf("Values after  swapping: %d %d", x, y)
}
