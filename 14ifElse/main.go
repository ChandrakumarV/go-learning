package main

func main() {

	x := 10
	if x > 10 {
		println("x > 10")
	} else if x == 10 {
		println("x == 10")
	} else {
		println("x < 10")
	}

	// decalare and assign in if statement
	if y := 20; y == 10 {
		println("y > 10", y)
	} else {
		println("y <= 10", y)
	}

	// println(y) only in if-else scope

}
