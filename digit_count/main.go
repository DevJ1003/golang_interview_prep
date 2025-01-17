package main

import (
	"fmt"
)

func main() {

	var num int
	var count = 0

	fmt.Print("Enter the number: ")
	fmt.Scan(&num)

	for num != 0 {

		num = num / 10
		count++
	}

	fmt.Printf("There are %d digits in given number!", count)
}
