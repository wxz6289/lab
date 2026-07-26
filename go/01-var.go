package main

import "fmt"

var PI = 3.14156926

const version = "1.0.0"

func var_declare() {
	var a int
	fmt.Println(a)
	var name string
	name = "King"
	fmt.Println(name)

	var age = 32
	fmt.Println(age)
	fmt.Printf("%T\n", age)

	// 函数内简短声明
	product := "Apple"
	fmt.Println(product)

	// 多变量声明
	var name1, age1 = "Dreamer", 26
	fmt.Println(name1, age1)

	fmt.Println(version)

	var (
		s1 string = "hello"
		s2 string = "go"
		b = true
	)

	fmt.Println(s1, s2, b)



}