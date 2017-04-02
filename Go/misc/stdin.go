package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewScanner(os.Stdin)
    reader.Scan()
    fmt.Printf(reader.Text())
    //text, _ := reader.ReadString('\n')
    //fmt.Println(text)
}
