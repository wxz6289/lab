package main

import (
	"fmt"
	"reflect"
)

func getType(obj any) {
	t := reflect.TypeOf(obj)
	fmt.Println(t, t.Kind())
	v := reflect.ValueOf(obj)
	fmt.Println(v, v.Int())
}

func show_type() {
	a := 12
	getType(a)
}