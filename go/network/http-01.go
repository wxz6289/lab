package main

import (
	"fmt"
	"net/http"
	"os"
)

func Index(writer http.ResponseWriter, request *http.Request){
	fmt.Println(request.URL.Path, request.UserAgent())
	byteData, err := os.ReadFile("./index.html")
	if err!=nil{
		return
	}
	writer.Write(byteData)
}

func main() {
	http.HandleFunc("/", Index)
	fmt.Println("Server running at 127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}