package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func error_show() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal("Fatal error:", err)
		panic("Panic error")
	} else {
		fmt.Println("Result:", result)
	}
}


func recover_show() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:", string(debug.Stack()))
		}
	}()
	value, err := divide(10, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("Value:", value)

	fmt.Println("This line will not be executed due to panic.")
}