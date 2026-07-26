package main

import "fmt"
func const_var() {
	const a = 12
	// a = 23
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	const (
		UP = iota  // iota仅能配合const使用
		DOWN
		LEFT
		RIGHT
	)

	fmt.Println(UP, DOWN, LEFT, RIGHT)

	const (
		b = 1*iota
		c
		d = 2*iota
		e
		f
	)

	// 0, 1, 4, 6, 8
	fmt.Println(b, c, d, e, f)

	const (
		X = 1 << iota
		W
		R
	)
	fmt.Println(X, W, R)

	const (
		A = iota
		B
		_
		C
		D = 5
		F
	)

	fmt.Println(A, B, C, D, F)
}