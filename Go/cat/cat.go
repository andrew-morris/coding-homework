package main

import "fmt"
import "os"

func main() {
	if len(os.Args) == 2 {
		argString := os.Args[1]
		fmt.Println(argString)
	}
}
