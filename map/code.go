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

}
