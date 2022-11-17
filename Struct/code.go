package main

import "fmt"

func main() {
	fmt.Println("Struct in Go language")
	// no inheritance support in golang : No super or parent interfaces
	avinash := User{"Avinash", "avinashlamichhane435@gmail.com", true, 16}
	fmt.Println(avinash)
	fmt.Printf("avinash details are: %+v\n", avinash)
	fmt.Printf("Name is %v\n", avinash.Name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
