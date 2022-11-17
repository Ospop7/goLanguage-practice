package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Go language")
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["Go"] = "Go"
	languages["Python"] = "Python"
	languages["Ruby"] = "Ruby"

	fmt.Println("List of all languages", languages)
	fmt.Println("JS shorts for", languages["JS"])

	// for deleting

	delete(languages, "Go")

	fmt.Println("List of all languages", languages)

	// loops in golang
	for k, v := range languages {
		fmt.Printf("For key %v, value is %v\n", k, v)
	}

	for _, v := range languages {
		fmt.Printf("For key , value is %v\n", v)
	}

}
