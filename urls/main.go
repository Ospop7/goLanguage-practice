package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to Handling the URLs in golang")

	fmt.Println(myurl)

	// parsing the url
	parsed, err := url.Parse(myurl)
	if err != nil {
		fmt.Println(err)

	}
	// fmt.Println(parsed.Host)
	// fmt.Println(parsed.Scheme)

	// fmt.Println(parsed.Path)

	// fmt.Println(parsed.Port())
	fmt.Println(parsed.RawQuery)

	qparams := parsed.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	// fmt.Println(qparams.Get("coursename"))
	fmt.Println(qparams["coursename"])

	// KEY VALUE PARSED SO WE APPLY THE FOR LOOP

	for _, v := range qparams {
		fmt.Println(v)

	}
	// for concat the url
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh"}

	// parsed, err = url.Parse(partsOfUrl.String())
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(parsed)

	// OR DO THIS

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}
