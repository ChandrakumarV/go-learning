package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Args Len: ", len(os.Args), os.Args)
}

// go run main.go test 123 d
