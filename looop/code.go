package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	fmt.Println(days)

	// for _, day := range days {
	// 	fmt.Println(day)
	// }

	// for d := 0; d < len(days); d++ {
	// 	// fmt.Println(days)
	// 	fmt.Println(days[d])

	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)

	// }

	value := 1

	for value < 10 {

		// if value == 4 {
		// 	break
		// }
		if value == 2 {
			goto lco
		}

		if value == 5 {
			value++
			continue
		}

		fmt.Println(value)
		value++

	}
lco:
	fmt.Println("jumped to")

}
