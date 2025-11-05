package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const basePath = "http://localhost:8000/"

func main() {
	getReq()
	postReq()
	formReq()

}

func getReq() {
	path := basePath

	res, err := http.Get(path)
	checkErr(err)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body) // body -> byte

	// fmt.Println(string(body))

	var responseString strings.Builder // struct assign like "new class()"
	responseString.Write(body)         // byte -> string.builder

	fmt.Println(responseString.String()) // string.builder -> string
}

func postReq() {

	path := basePath + "post"
	requestBody := strings.NewReader(`
		{
			"name":"chandru",
			"age":13,
			"isStudent":true
		}
	`)

	res, err := http.Post(path, "application/json ", requestBody)
	checkErr(err)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))

}

func formReq() {

	path := basePath + "postform"

	data := url.Values{}

	data.Add("name", "chandru")
	data.Add("age", "12")
	data.Add("isStudent", "true")

	res, err := http.PostForm(path, data)
	checkErr(err)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
