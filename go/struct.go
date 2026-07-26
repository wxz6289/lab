package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
	Password string `json:"-"`
}

type Student struct {
	Person
	Grade string `json:"grade"`
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func (s Student) Info() {
	fmt.Printf("Hello, my name is %s, I am %d years old and I am in grade %s.\n", s.Name, s.Age, s.Grade)
}

func (p *Person) SetName(name string) {
	fmt.Printf("After changing : %p\n", p)
	p.Name = name
}

func (p Person) SetAge(age int) {
	fmt.Printf("After changing age : %p\n", &p)
	p.Age = age
}

func struct_show() {
	person := Person{Name: "Alice", Password: "secret"}
	person.Greet()

	student := Student{Person: Person{Name: "Bob", Age: 20, Password: "secret"}, Grade: "A"}
	student.Greet()
	student.Info()
	fmt.Printf("Before changing : %p\n", &person)
	person.SetAge(35)
	person.SetName("Charlie")
	person.Greet()
	bytes, _ := json.Marshal(person)
	fmt.Println(string(bytes))
}