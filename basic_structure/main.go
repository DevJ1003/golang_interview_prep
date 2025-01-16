// basic structure of a golang program

// package defines the package
package main

// import used to import third party and external libraries
import (
	"fmt"
	"math"
)

// we have to define global variable and constant & struct outside main function
// short names should start from capital letter
// long names should be in camelcase form
const Pi = 3.14

var golabalVar = "I am a global variable"

type Person struct {
	Name string
	Age  int
}

// the main function is the first function which loads in our program
func main() {

	fmt.Println("Hello world!")

	person := Person{
		Name: "test",
		Age:  22,
	}

	fmt.Println(golabalVar)
	fmt.Printf("The name is %s and age is %d\n", person.Name, person.Age)
	fmt.Printf("Square root of 16 is %.2f\n", math.Sqrt(16))
}
