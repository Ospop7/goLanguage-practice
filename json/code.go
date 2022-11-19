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
	EncodeJson()

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
