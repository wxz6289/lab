package main

import "fmt"

func show_types() {
	var i int = 42
	var f float64 = 3.14159
	var b bool = true
	var s string = "Hello, Go!"

	var a rune = '法'

	var ms = `
	This is a multi-line string.

	It can span multiple lines.
	`

	fmt.Printf("i: %v (%T)\n", i, i)
	fmt.Printf("f: %v (%T)\n", f, f)
	fmt.Printf("b: %v (%T)\n", b, b)
	fmt.Printf("s: %v (%T)\n", s, s)
	fmt.Printf("a: %v (%T)\n", a, a)
	fmt.Printf("ms: %v (%T)\n", ms, ms)
}

type Code int
type CodeAlias = int

func type_show() {
	var c Code = 42
	var ca CodeAlias = 42

	fmt.Printf("c: %v (%T)\n", c, c)
	fmt.Printf("ca: %v (%T)\n", ca, ca)
}