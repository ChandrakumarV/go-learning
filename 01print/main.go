package main

import (
	"fmt"
)

func PrintFuncs() {

	const a, b = "test", 123

	fmt.Println(a, b)
	fmt.Print("Print")
	fmt.Println("Println")
	fmt.Printf("Verbs   :  %v, %T, %d, %.2f, %s, %q \n", 3, 3, 4, 4.555, "hello", "hello")
	fmt.Printf("%x %X %U \n", "a", "A", 'c')

	// Variable declaration and assignment
	var c bool
	c = true

	// Short hand declaration and assignment
	// only inside a function
	d := "short"
	fmt.Println(d)

	// Multiple variable declaration and assignment
	var x, y int = 1, 2
	fmt.Println(x, y)

	// Multiple short hand declaration and assignment
	m, n := 3, 4
	fmt.Println(m, n)

	// Blank identifier
	_, z := 5, 6
	fmt.Println(z)

	// Constants
	const pi = 3.14
	fmt.Println(pi)

	// Multiple constant declaration
	const (
		e   = 2.71
		g   = 9.8
		sol = "Sun"
	)
	fmt.Println(e, g, sol)

	// iota - enumerated constants
	const (
		cat   = iota // 0
		dog          // 1
		lion         // 2
		tiger        // 3
	)
	fmt.Println(cat, dog, lion, tiger) // 0 1 2 3

	// iota with expressions
	const (
		a1 = iota * 10 // 0 * 10 = 0
		b1             // 1 * 10 = 10
		c1             // 2 * 10 = 20
	)
	fmt.Println(a1, b1, c1) // 0 10 20

	// iota reset
	const (
		x1 = iota // 0
		x2        // 1
		x3        // 2
		x4        // 3
	)
	fmt.Println(x1, x2, x3, x4) // 0 1 2 3

	// iota with multiple constants
	const (
		y1, y2 = iota + 1, iota + 2 // 0+1=1, 0+2=2
		z1, z2                      // 1+1=2, 1+2=3
		w1, w2                      // 2+1=3, 2+2=4
	)
	fmt.Println(y1, y2, z1, z2, w1, w2) // 1 2 2 3 3 4

	// iota with skipping
	const (
		p1 = iota // 0
		_         // skip 1
		p3 = iota // 2
	)
	fmt.Println(p1, p3) // 0 2
}
