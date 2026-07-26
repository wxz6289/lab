package main

import "fmt"

func run() {
    fmt.Printf("%T\n", 2.1)
    fmt.Printf("%#v\n", "Hello, World!")

    var s = fmt.Sprintf("%s, %d, %.2f", "Hello", 2+3, PI)
    fmt.Println(s)
}

func input() {
    fmt.Println("Please input your name:")
    var name string
    fmt.Scanln(&name)
    fmt.Println("Hello", name)
}