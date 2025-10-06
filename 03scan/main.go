package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a number :")

	input, _ := reader.ReadString('\n')

	// num, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // for float
	num, err := strconv.ParseInt(strings.TrimSpace(input), 10, 32) //for int

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("\n Entered no is ", num)
	}
}
