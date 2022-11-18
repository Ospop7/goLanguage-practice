package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)
	// proRes, _ := proAdder(2, 3, 4, 5, 6) for ignore the seconf value
	proRes, myMessage := proAdder(2, 4, 35, 5)

	fmt.Println("Result is: ", proRes)

	fmt.Println("Pro Result is: ", proRes)
	fmt.Println("Pro Message is: ", myMessage)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo

}
func greeterTwo() {
	fmt.Println("Another function called")

}
func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}

	return total, "Hello Pro result functions"

}

func greeter() {
	fmt.Println("functions in golang")

}
