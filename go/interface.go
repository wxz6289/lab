package main

import "fmt"

type Greeter interface {
	Hello() string
}


func (p Person) Hello() string {
	return "Hello, my name is " + p.Name
}

func greet(g Greeter) {
	// 类型断言
	p := g.(Person)
	fmt.Println(p.Hello())
}

	type Animal struct {
		Name string
	}

	func (a Animal) Hello() string {
		return "Hello, I am an animal named " + a.Name
	}

type Print interface {
}

// type any = interface{}

func print(p any) {
	fmt.Println(p)
}

func interface_show() {
	p := Person{Name: "Alice"}
	greet(p)
	var g Greeter

	g = p
	fmt.Println(g.Hello())

	a := Animal{Name: "Buddy"}
	fmt.Println(a.Hello())
	// greet(a)
	print(a)
}

