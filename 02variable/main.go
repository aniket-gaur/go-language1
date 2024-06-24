package main

import "fmt"

//public variable

var Logintoken string = "12345566"

func main() {
	// strings in go

	var username string = "Aniket"
	fmt.Println(username)
	// to get the variable of type
	//Printf is format specifier used for standard output

	fmt.Printf("Variable is of type:%T \n ", username)

	// boolean type

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type:%T \n", isLoggedIn)

	// int types various on basics of range
	// unit8 means unsigned bits
	//where as int8 stands for signed bits
	var a uint8 = 24
	fmt.Println(a)
	fmt.Printf("Variable is of type:%T \n", a)

	// float values
	var b float32 = 24.345
	fmt.Println(b)
	fmt.Printf("Variable is of type:%T \n", b)

	// imp implict types
	//if not specified the varible the lexer will decide on the basis of value declared

	var website = "google.com"

	fmt.Println(website)

	//no var keyword

	//we can use this operator inside methods not globally

	c := 34
	fmt.Println(c)
	fmt.Println(Logintoken)

	// some imp features of go
	// 1. same line declarations

	mileage, car := "123", "Honda City"

	fmt.Println(mileage, car)

	// 2. return a formatted string
	//how to use

	s := fmt.Sprintf("I am %v years old", 20)

	fmt.Println(s)

}
