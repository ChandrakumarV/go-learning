package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1") // JSON - response
	// resp, err := http.Get("http://example.com/") // HTML - Response

	checkNilErr(err)

	// must close the request even write or read the response/request
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // this will give the byte output

	checkNilErr(err)

	fmt.Println(string(body)) //byte to string

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
