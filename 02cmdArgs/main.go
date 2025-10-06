package main

import (
	"fmt"
	"os"
)

func CmdArgs() {
	fmt.Println("Args Len: ", len(os.Args), os.Args)
}
