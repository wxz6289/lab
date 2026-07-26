package main

import (
	"fmt"
	"sort"
)

func array_show() {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	var arr2 = [...]int{6, 7, 8, 9, 10}

	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)

	for i := 0; i < len(arr1); i++ {
		fmt.Println("arr1[", i, "] =", arr1[i])
	}

	for i := range arr2 {
		fmt.Println("arr2[", i, "] =", arr2[i])
	}
}

func slice_show() {
	var str []string
	str = append(str, "Hello")
	str = append(str, "World")
	fmt.Println(str)

	str2 := []string{"Go", "is", "awesome"}
	fmt.Println(str2)

	str3 := make([]int, 5)
	for i := 0; i < len(str3); i++ {
		str3[i] = i * 2
	}
	fmt.Println(str3)

	str4 := str2[1:3]
	fmt.Println(str4)
}

func slice_sort() {
	var numbers = []int{5, 2, 8, 1, 4}
	fmt.Println("Before sorting:", numbers)

	sort.Ints(numbers)
	fmt.Println("After sorting:", numbers)
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println("After reversing:", numbers)
}
	// Sort the slice in ascending order