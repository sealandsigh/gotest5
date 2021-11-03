package main

import (
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

func main() {
	http.HandleFunc("/test/", f1)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
