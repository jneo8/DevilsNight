package main

import "fmt"

func main() {
    var input string

    fmt.Printf("Enter your name: ")
    fmt.Scanln(&input)
    fmt.Printf("Hello, %s!\n", input)
}
