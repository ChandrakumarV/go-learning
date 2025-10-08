package main

import (
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(6) + 1

	switch num {
	case 1:
		println("One")
	case 2:
		println("Two")
		fallthrough // fall through will continue to the next case even if it doesn't match
	case 3:
		println("Three")
		fallthrough
	case 4:
		println("Four")
	case 5:
		println("Five")
	default:
		println("Six")
	}
}
