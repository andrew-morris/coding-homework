package main

import (
    "fmt"
    "os"
	"strings"
	"io/ioutil"
)


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func grepFile(nameOfFileToRead string, keywordToLookFor string) {
    data, err := ioutil.ReadFile(nameOfFileToRead)
    check(err)
	s := string(data)
	slicedFile := strings.Split(s, "\n")
	for _, x := range slicedFile {
		if strings.Contains(x, keywordToLookFor) {
            fmt.Println(x)
        }
	}
}

func main() {
    if len(os.Args) == 3 {
        argFilename := os.Args[2]
        argKeyword := os.Args[1]
        grepFile(argFilename, argKeyword)
    }
}
