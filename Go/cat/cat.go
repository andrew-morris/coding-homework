package main

import "fmt"
import "os"
import "io/ioutil"

//TODO: Add support for multiple files
//TOOD: Less disgusting errors

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(nameOfFileToRead string) {
	data, err := ioutil.ReadFile(nameOfFileToRead)
	check(err)
	fmt.Print(string(data))
}

func main() {
	if len(os.Args) == 2 {
		argString := os.Args[1]
		readFile(argString)
	}
}
