// basic structure of a golang program

package main

import (
	"fmt"
	"math"
)

const Pi = 3.14

var golabalVar = "I am a global variable"

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Hello world!")

	person := Person{
		Name: "test",
		Age:  22,
	}

	fmt.Printf("The name is %s and age is %d\n", person.Name, person.Age)
	fmt.Printf("Square root of 16 is %.2f\n", math.Sqrt(16))
}
