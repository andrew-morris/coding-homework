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

func tailFile(nameOfFileToRead string) {
    data, err := ioutil.ReadFile(nameOfFileToRead)
    check(err)
    s := string(data)
    slicedFile := strings.Split(s, "\n")
    lastTen := slicedFile[len(slicedFile)-10:]
    for _, x := range lastTen {
        fmt.Println(x)
    }
}

func main() {
    if len(os.Args) == 2 {
        argFilename := os.Args[1]
        tailFile(argFilename)
    } else {
        fmt.Println("tail <filename>")
        fmt.Println("tail /etc/passwd")
    }
}
