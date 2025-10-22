package main

import (
	"fmt"
	"net/url"
)

var myUrl = "https://www.youtube.com/watch?name=chandru&age=21&course=science&course=sdfsdfdsf"

func main() {

	// Destructure URL -------------

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)   // https
	fmt.Println(result.Host)     // www.youtube.com
	fmt.Println(result.Path)     // /watch
	fmt.Println(result.Port())   // ""
	fmt.Println(result.RawQuery) // name=chandru&age=21&course=science

	// Map search querys -----------

	mapQuery := result.Query()

	for key, value := range mapQuery {
		fmt.Println(key, " : ", value)
	}

	// Build URL -------------------

	partsOfURL := &url.URL{
		Scheme:   "https",
		Host:     "www.test.com",
		Path:     "/check",
		RawQuery: "user=vicky",
	}

	fmt.Println()
	fmt.Println(partsOfURL.String())

}
