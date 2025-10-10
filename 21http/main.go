package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://example.com/")
	checkNilErr(err)

	// must close the request even write or read the response/request
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	checkNilErr(err)

	fmt.Println(string(body))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
