package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http server

func f1(w http.ResponseWriter, r *http.Request) {
	str, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		w.Write([]byte("暂无内容"))
	}
	w.Write([]byte(str))
}

func f2(w http.ResponseWriter, r *http.Request) {
	// 对于GET请求，参数都放在url上，请求体是没有数据的
	queryParam := r.URL.Query()
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age) //自动帮我们识别url中的queary param
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body)) // 我在服务端打印客户端来的body
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/test/", f1)
	http.HandleFunc("/client/", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
