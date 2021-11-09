package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// http_client

// 共用一个client适用于 请求比较频繁
var (
	client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
		},
	}
)

func main() {
	// resp, err := http.Get("http://127.0.0.1:9090/client/?name=sb&age=18")
	// if err != nil {
	// 	fmt.Println("get url failed,err:", err)
	// 	return
	// }
	data := url.Values{} // url values
	urlObj, _ := url.Parse("http://127.0.0.1:9090/client/")
	data.Set("name", "周林")
	data.Set("age", "9000")
	queryStr := data.Encode() // URL encode之后的URL
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	if err != nil {
		fmt.Println("NewRequest failed,err:", err)
		return
	}
	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Println("DefaultClient failed,err:", err)
	// 	return
	// }
	// defer resp.Body.Close() // 一定要记得关闭

	// 请求不是特别频繁，用完关闭该链接
	// 禁用keepalive的client
	// tr := &http.Transport{
	// 	DisableKeepAlives: true,
	// }
	// client := http.Client{
	// 	Transport: tr,
	// }
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("DefaultClient failed,err:", err)
		return
	}
	defer resp.Body.Close() // 一定要记得关闭

	// 发请求
	// 从resp中把服务端返回的数据读出来
	// var data []byte
	// resp.Body.Read()
	// resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body) // 我在客户端读出server响应的body
	if err != nil {
		fmt.Println("read resp.body failed,err:", err)
		return
	}
	fmt.Println(string(b))
}
