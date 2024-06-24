package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is time module ")
	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdTime := time.Date(2024, time.July, 24, 12, 13, 19, 0, time.Local)

	fmt.Println(createdTime)

	fmt.Println(createdTime.Format("01-02-2006 15:04:05 Monday"))
}
