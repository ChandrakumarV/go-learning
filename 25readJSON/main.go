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
	jsonDataFormWeb := []byte(`
		  {
            "courseName": "JS",
            "Price": 199,
            "Platform": "Online",
            "tags": [
            	"JS",
                "Developer",
                "Frontend"
            ]
          }
		`)

	var course Course

	checkValid := json.Valid(jsonDataFormWeb)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFormWeb, &course)
		fmt.Printf("%#v\n", course)
		// fmt.Println(course)
	} else {

		fmt.Println("JSON is not valid")
	}

	// In some case we just need to add the extra key without the struct

	fmt.Println("-------------------- Without Struct")
	var onlineCourse map[string]interface{}
	json.Unmarshal(jsonDataFormWeb, &onlineCourse)
	fmt.Printf("%#v\n", onlineCourse)

}
