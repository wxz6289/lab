package main

import "fmt"

func map_show() {
	var m = make(map[string]int)
	m["apple"] = 5
	m["banana"] = 10
	fmt.Println(m)

	v, ok := m["apple"]
	if ok {
		fmt.Println("Value for 'apple':", v)
	}

	delete(m, "banana")
	fmt.Println(m)

	/* var m2 map[string]string
	m2["king"] = "dreamer" // This will cause a runtime panic because m2 is nil */

	m3 := make(map[string]string)
	m3["king"] = "dreamer"
	fmt.Println(m3)
}