package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Welcome to vedio on slices")

	// var fruitList = []string{"Apple ", "Tomato", "Peach"}
	// fmt.Printf("Type of fruitList is %T\n", fruitList)

	// fruitList = append(fruitList, "banana", "Mango")
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	// // dynamic memory allocations

	// highScore := make([]int, 4)

	// highScore[0] = 234
	// highScore[1] = 23
	// highScore[2] = 34
	// highScore[3] = 24

	// highScore = append(highScore, 45, 55, 232, 55)

	// fmt.Println(highScore)

	// // for sorting an int value
	// fmt.Println(sort.IntsAreSorted(highScore))

	// sort.Ints(highScore)
	// fmt.Println(highScore)
	// fmt.Println(sort.IntsAreSorted(highScore))

	// how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "typescript", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
