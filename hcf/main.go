package main

import (
	"fmt"
)

func main() {

	var x, y, hcf int
	var i = 1

	fmt.Print("Enter the values for x & y: ")
	fmt.Scan(&x, &y)

	for i = 1; i <= x && i <= y; i++ {
		if x%i == 0 && y%i == 0 {
			hcf = i
		}
	}

	fmt.Printf("The HCF for %d & %d is %d", x, y, hcf)
}
