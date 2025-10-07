package main

import "fmt"

// go don't have inheritance, super or base

type Person struct {
	name     string
	age      int
	isActive bool
}

func main() {

	// Method 1 -----------
	s := Person{}
	s.name = "Chandru"
	s.age = 30
	s.isActive = true

	fmt.Println(s)

	// Method 2 ---------------
	s2 := Person{name: "Chandru", age: 30, isActive: true}
	fmt.Println(s2)

	// Method 3 -----------
	s3 := Person{"Chandru", 30, true}
	fmt.Println(s3)

	// Empty struct
	var s4 struct{}
	fmt.Printf("%#v\n", s4)

}
