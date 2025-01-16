package main

import (
	"fmt"
)

// Explicit declaration with the "var" keyword =>  You specify the type of the variable.
// Short variable declaration using ":=" => The type is inferred based on the value assigned.

func main() {
	// 1. Declaration with `var` keyword (explicit type)
	var age int = 25
	var name string = "Alice"
	var isActive bool = true

	// 2. Declaration with `var` keyword (type inferred)
	var city = "New York"
	var temperature = 22.5

	// 3. Short variable declaration using `:=` (type inferred)
	country := "USA"
	population := 331_000_000
	isSunny := false

	// Print all the variables
	fmt.Println("Explicit declaration:")
	fmt.Println("Age:", age, "Name:", name, "Is Active:", isActive)

	fmt.Println("\nType inferred with `var`:")
	fmt.Println("City:", city, "Temperature:", temperature)

	fmt.Println("\nShort variable declaration:")
	fmt.Println("Country:", country, "Population:", population, "Is Sunny:", isSunny)
}
