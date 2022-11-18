package main

import "fmt"

func main() {
	// last In first out
	defer fmt.Println("World!")
	defer fmt.Println("One World!")
	fmt.Println("hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)

	}
}
