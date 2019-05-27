package main

import (
	"fmt"
	"os"

	"github.com/rahulvramesh/kitago-v2/cmd"
)

func main() {

	//only min 2, max 4 input

	if len(os.Args) < 3 || len(os.Args) > 5 {
		fmt.Println("argument should be 2 or 4 ")
		os.Exit(1)
	}

	fmt.Println(cmd.ExecByName(string(os.Args[1]), os.Args))
}
