package main

import "fmt"

// go don't have class,so we'll use the method in struct

type Person struct {
	name     string
	age      int
	isActive bool
	oneAge   int
}

func (p Person) GetStatus() {
	fmt.Println("Is user Active?: ", p.isActive)
}

func (p Person) UpdateName() {
	p.name = "new name"
	fmt.Println("Updated name: ", p.name)
}

func main() {

	p := Person{name: "chandru", age: 24, isActive: true}

	fmt.Printf("%+v\n", p)
	p.GetStatus()
	p.UpdateName()
	fmt.Printf("%+v\n", p) // name will not be updated because in method we are using value receiver
	// to update the name we need to use pointer receiver

}
