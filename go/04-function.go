package main

import "fmt"

func sum(nums ...int) int {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
}


func function_show() {
		fmt.Println(sum(1, 2, 3, 4, 5))
}


func return_multiple_values() (val int, ok bool) {
	val = 3
	ok = true
	return
}

func return_multiple_values_show() {
		val, ok := return_multiple_values()
		fmt.Println("val:", val, "ok:", ok)
}


func void_name() {
		var test_fn = func () {
			fmt.Println("This is a void function.")
		}
		test_fn()
}

/* func init() {
	fmt.Println("This is the init function. in function.go")
} */

func defer_show() {
		defer fmt.Println("This is the deferred function.")
		name := "John"
		defer func() {
			fmt.Println("This is the deferred anonymous function.", name)
		}()
		age := 30
		defer func(age int) {
			fmt.Println("This is the deferred anonymous function with parameters.", age)
		}(age)
		defer fmt.Println("This is the deferred function2.")
		fmt.Println("This is the main function.")
}

func name_return() (a, b int, f string) {
	a = 12
	f = "King"
	return
}

func show_name_return() {
	e, f, g := name_return()
	fmt.Println(e, f, g)
}