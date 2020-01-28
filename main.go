package main

import (
	"fmt"
)

type astrosDataFormat struct {
	People  []peopleDataFormat `json:"people"`
	Number  int                `json:"number"`
	Message string             `json:"message"`
}
type peopleDataFormat struct {
	Name  string `json:"name"`
	Craft string `json:"craft"`
}

func main() {

	url := "http://api.open-notify.org/astros.json"
	json := Get(url)

	astros := astrosDataFormat{}
	Decode(string(json), &astros)

	fmt.Printf("\n\n json object:::: %+v", astros)
}
