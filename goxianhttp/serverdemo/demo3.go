package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/get/hello", hello)
	fmt.Println("0.0.0.0:8085/get/hello")

	http.ListenAndServe(":8085", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if len(query["name"]) > 0 {
		names := query["name"][1]
		fmt.Println("字典下标的值", names)
	}
	fmt.Println(query)

	w.Write([]byte("hello world"))
}
