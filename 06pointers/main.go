package main

import "fmt"

func main() {

	fmt.Println("this is pointers module ")

	myNumber := 24
	var ptr = &myNumber

	fmt.Println("value of actual pointer", ptr)
	fmt.Println("value of actual variable", *ptr)
}
