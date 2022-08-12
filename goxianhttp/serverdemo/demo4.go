package main

import (
	"encoding/json"
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

	name2 := query.Get("name")
	fmt.Println("Get方法获取的值", name2)

	type Info struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Age      int    `json:"age"`
	}

	u := Info{Name: name2, Password: "123456", Age: 26}
	json.NewEncoder(w).Encode(u)
	fmt.Println(query)

	w.Write([]byte("hello world"))
}
