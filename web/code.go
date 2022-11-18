package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	fmt.Println("LOC web request")

	resp, err := http.Get(url)
	if err != nil {
		panic(err)

	}
	fmt.Printf("Respone is of type %T\n", resp)

	defer resp.Body.Close() // caller's reponsibility to close the connections

	databytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)

	}
	content := string(databytes)
	fmt.Printf(content)

}
