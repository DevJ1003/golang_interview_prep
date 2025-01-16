package main

import (
	"fmt"
)

func main() {
	// Declare variables without initialization
	var intVar int
	var floatVar float64
	var boolVar bool
	var stringVar string
	var sliceVar []int
	var mapVar map[string]int
	var pointerVar *int
	var funcVar func()
	var structVar struct {
		name string
		age  int
	}

	// Print zero values
	fmt.Println("Zero value examples in Go:")
	fmt.Println("int:", intVar)
	fmt.Println("float64:", floatVar)
	fmt.Println("bool:", boolVar)
	fmt.Println("string:", stringVar)
	fmt.Println("slice:", sliceVar)
	fmt.Println("map:", mapVar)
	fmt.Println("pointer:", pointerVar)
	fmt.Println("function:", funcVar)
	fmt.Println("struct:", structVar)
}
