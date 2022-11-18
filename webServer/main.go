package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	fmt.Printf("Welcome to webserver!\n")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"

	respo, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer respo.Body.Close()

	fmt.Println("Status: ", respo.StatusCode)
	fmt.Println("Content length is ", respo.ContentLength)
	var responseString strings.Builder

	content, _ := ioutil.ReadAll(respo.Body)

	byteCount, _ := responseString.Write(content)
	fmt.Println("Bytes: ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(content) is in bytes forms
	// fmt.Println(string(content))

}
