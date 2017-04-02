package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func headFile(nameOfFileToRead string) {
	data, err := ioutil.ReadFile(nameOfFileToRead)
	check(err)
	s := string(data)
	slicedFile := strings.Split(s, "\n")
	firstTen := slicedFile[0:10]
	for _, x := range firstTen {
		fmt.Println(x)
	}
}

func main() {
	if len(os.Args) == 2 {
		argFilename := os.Args[1]
		headFile(argFilename)
	} else {
		fmt.Println("head <filename>")
		fmt.Println("head /etc/passwd")
	}
}
