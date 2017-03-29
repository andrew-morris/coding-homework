package main

import "fmt"
import "os"

func main() {
	fmt.Printf("Args: %d\n\n", len(os.Args))
	for i, e := range os.Args {
		fmt.Printf(" %d: %s\n", i, e)
	}
	fmt.Println()
}
