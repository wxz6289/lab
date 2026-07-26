package main

import "fmt"

/* func init() {
	fmt.Println("This is the init function. in closure.go")
} */

func counter() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
}

func closure_show() {
	counter1 := counter()
	fmt.Println(counter1()) // 输出: 1
	fmt.Println(counter1()) // 输出: 2
	fmt.Println(counter1()) // 输出: 3

	counter2 := counter()
	fmt.Println(counter2()) // 输出: 1
	fmt.Println(counter2()) // 输出: 2
	fmt.Println(counter2()) // 输出: 3
}