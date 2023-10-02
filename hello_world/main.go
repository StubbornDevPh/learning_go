package main

import "fmt"

func main() {
	// print string
	fmt.Println("Hello world")
	// declaring an int variable
	var testvar int = 1
	// declaring a pointer
	var pTestVar *int = &testvar
	// printing pointer memory location
	fmt.Println(pTestVar)
	// printing pointer value
	fmt.Println(*pTestVar)

	// Type declarations like a boss
	type numberType int
	type booleanType bool

	var numTest numberType = 1
	var isTrue booleanType = true
	fmt.Println(numTest, isTrue)

	// one thing that leverages type declaration is type conversion
	var numbahComputation = numberType((1 + 1))
	fmt.Println(numbahComputation)

}
