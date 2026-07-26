package main

import (
	"fmt"
	"net/http"
)
func main() {
	response, err := http.Get("http://127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	
}