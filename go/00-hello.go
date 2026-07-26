package main

import (
	"fmt"

	ver "github.com/wxz6289/go/version"

	"rsc.io/quote/v4"
)

func hello(){
	var a = 2 + 3
	var name string = "King"
	fmt.Println(name)
	fmt.Println("Hello Go!", a)
	var str = fmt.Sprintf("%s, %d", "Hello", a)
	fmt.Println(str)
	fmt.Println(quote.Go())
	var_declare()
	fmt.Println(PI)
	fmt.Println(ver.Version)
	fmt.Println(version)
}

func main() {
	// hello()
	// var_declare()
	// const_var()
	show_name_return()
/*
	run()
	input() */
	// show_types()
	// array_show()
	// slice_show()
	// slice_sort()
	// map_show()
	// if_else()
	// switch_case()
	// for_loop()
	// dead_loop()
	// iterate_map()
	// closure_show()
	// function_show()
	// return_multiple_values_show()
	// void_name()
	// defer_show()
	// struct_show()
	// type_show()
	// interface_show()
	// show()
	// show_goroutine()
	// show_pay()
	// show_send()
	// show_see()
	// show_total()
	// show_map()
	// safe_map_show()
	// error_show()
	// recover_show()
	// show_generic_type()
	// show_file()
	// show_file_stream()
	// show_file_line()
	// show_file_by_scanner()
	// show_write_file()
	// show_append_file()
	// show_copy_file()
	// show_list_files_in_directory()
	// show_type()
	// server()
	// client()
}

