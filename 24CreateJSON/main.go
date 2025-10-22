package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string `json:"courseName"` // rename key in JSON
	Price    int
	Platform string
	Password string   `json:"-"`              // remove from JSON payload
	Tags     []string `json:"tags,omitempty"` // remove the property if null/nil value
}

func main() {
	courses := []Course{
		{
			"JS",
			199,
			"Online",
			"zawezxas",
			[]string{"JS", "Developer", "Frontend"},
		},
		{
			"GO",
			449,
			"Offline",
			"zawezxas",
			nil, // look this in json
		},
	}

	// finalJson, err := json.Marshal(courses)
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	checkErr(err)

	fmt.Println("JSON ------------------")
	fmt.Println(string(finalJson))

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
