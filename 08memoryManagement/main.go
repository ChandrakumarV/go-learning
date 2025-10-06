package main

import (
	"fmt"
	"runtime"
)

// GOGC=100 | by default
// GOGC=off | if needed

func main() {
	cpuCount := runtime.NumCPU()
	fmt.Println(cpuCount)

}
