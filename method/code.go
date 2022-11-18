package main

import "fmt"

func main() {
	fmt.Println("Struct in Go language")
	// no inheritance support in golang : No super or parent interfaces
	avinash := User{"Avinash", "avinashlamichhane435@gmail.com", true, 16}
	fmt.Println(avinash)
	fmt.Printf("avinash details are: %+v\n", avinash)
	fmt.Printf("Name is %v\n", avinash.Name)
	fmt.Printf("Email is %v\n", avinash.Email)

	avinash.GetStatus()
	avinash.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "avinah@gmail.com"
	fmt.Println("Email of this user: ", u.Email)
}
