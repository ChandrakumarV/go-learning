package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	writeFile("file.txt")
	readFile("file.txt")

}

func writeFile(fileName string) {
	content := "Chandru is good boy"

	file, err := os.Create(fileName) //create file
	checkNilErr(err)

	defer file.Close()

	len, err := io.WriteString(file, content) //write to file
	checkNilErr(err)

	fmt.Println("Wrote Length is : ", len)
	fmt.Println("File created successfully")
}

func readFile(fileName string) {

	content, err := os.ReadFile(fileName) //read file
	checkNilErr(err)
	fmt.Println("content : ", string(content))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
