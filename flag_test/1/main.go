package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
	// go build main.go
	//// ./main.exe a b c d e f
}
