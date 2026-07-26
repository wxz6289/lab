package main

import (
	"fmt"
	"time"
)

func if_else() {
	var a = 10
	if a > 5 {
		fmt.Println("a is greater than 5")
	} else {
		fmt.Println("a is less than or equal to 5")
	}

	if b := a * 2; b > 15 {
		fmt.Println("b is greater than 15")
	} else {
		fmt.Println("b is less than or equal to 15")
	}
}

func switch_case() {
	// 默认不需要使用 break 语句，Go 的 switch 会自动在每个 case 后面加上 break，需要使用 fallthrough 关键字来实现贯穿效果
	var day = 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Another day")
	}

	switch {
	case day < 5:
		fmt.Println("It's a weekday")
	default:
		fmt.Println("It's a weekend")
	}
}

func for_loop() {
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	// 模拟 while 循环
	j := 0
	for j < 5 {
		fmt.Println("j:", j)
		j++
	}

	k := 0
	for {
		if k >= 5 {
			break
		}
		fmt.Println("k:", k)
		k++
	}
}

func dead_loop() {
	fmt.Println("This is an infinite loop. Press Ctrl+C to stop.")
	for {
		fmt.Println(time.Now())
		time.Sleep(1 * time.Second)
	}
}

func iterate_map() {
	var myMap = map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 7,
	}

	for key, value := range myMap {
		fmt.Printf("%s: %d\n", key, value)
	}

	array := []int{1, 2, 3, 4, 5}
	for index, value := range array {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}