package main

import "fmt"

type Number interface {
	int | float64
}


func sum2[T Number](a, b T) T {
	return a + b
}

func show_generic_type() {
	var intResult int = sum2(3, 5)
	var floatResult float64 = sum2(3.5, 2.521)

	println("intResult:", intResult)
	fmt.Printf("floatResult:%.2f\n", floatResult)
}