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
			Name:     "JS",
			Price:    199,
			Platform: "Online",
			Password: "zawezxas",
			Tags:     []string{"JS", "Developer", "Frontend"},
		},
		{
			"GO",
			449,
			"Offline",
			"zawezxas",
			nil, // look this in json
		},
	}

	// finalJson, err := json.Marshal(courses) // without aligment
	finalJson, err := json.MarshalIndent(courses, "", "\t") // with aligment

	checkErr(err)

	fmt.Println("JSON ------------------")
	fmt.Println(string(finalJson))

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
