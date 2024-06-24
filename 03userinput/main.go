package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "this is userinput module "

	fmt.Println(welcome)

	// reading from user using library

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the Rating pls:")

	//the input from the user needs to be saved
	//comma ok || err synatx
	//in case we want to ignore err use _ underscore in place of err

	input, err := reader.ReadString('\n')

	fmt.Println("Thanks For the Rating", input)
	fmt.Printf("The type of Rating is %T", input)
	fmt.Println("the Error is ", err)

}
