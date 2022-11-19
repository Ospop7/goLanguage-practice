package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	platfrom string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags, omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")
	// EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "avdefbfr", []string{"web-dev", "js"}},
		{"AngularJS Bootcamp", 29, "LearnCodeOnline.in", "avdeew444r", []string{"web-dev", "js"}},

		{"Golang Bootcamp", 99, "LearnCodeOnline.in", "avd9090", nil},
	}

	// package this data as a json data

	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

// consume JSON data in golang
func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "AngularJS Bootcamp",
		"Price": 29,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
}
 `)

	var lcoCourses course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("valid JSON")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)

	} else {
		fmt.Println("invalid JSON")

	}

	// some cases where you just want to add data to key value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {

		fmt.Printf("Key is %s and value isn %v\n", k, v)
	}

}
