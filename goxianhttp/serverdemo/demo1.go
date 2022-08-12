package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/get/hello", hello)

	http.ListenAndServe(":8085", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println(query)

	w.Write([]byte("hello world"))
}
