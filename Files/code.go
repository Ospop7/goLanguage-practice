package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a files"

	file, err := os.Create("./mygofile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("length is :", length)

	defer file.Close()
	readFile("./mygofile.txt")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Text is: \n", string(databyte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
