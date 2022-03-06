package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// gin的helloWorld

//func main() {
//	// 1. 创建路由
//	// 默认使用了2个中间件Logger(), Recovery()
//	r := gin.Default()
//	// 也可以创建不带中间件的路由
//	// r:= gin.New()
//	// 2. 绑定路由规则,执行函数
//	// gin.Context,封装了 request和response
//	r.GET("/", func(c *gin.Context) {
//		c.String(http.StatusOK, "hello world!")
//	})
//	// 3. 监听端口，默认在8080
//	err := r.Run(":8100")
//	if err != nil {
//		return
//	}
//}

//func main() {
//	r := gin.Default()
//	// api参数
//	r.GET("/user/:name/*action", func(c *gin.Context) {
//		name := c.Param("name")
//		action := c.Param("action")
//		c.String(http.StatusOK, name+" is "+action)
//	})
//	err := r.Run(":8100")
//	if err != nil {
//		return
//	}
//}
//
//func main() {
//	r := gin.Default()
//	// api参数
//	r.GET("/welcom", func(c *gin.Context) {
//		// DefaultQuery 第二个参数是默认值
//		name := c.DefaultQuery("name", "jack")
//		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
//	})
//	err := r.Run(":8100")
//	if err != nil {
//		return
//	}
//}

//func main() {
//	r := gin.Default()
//	// url参数
//	r.POST("/form", func(c *gin.Context) {
//		// 表单参数设置默认值
//		type1 := c.DefaultPostForm("type", "alert")
//		// 接收其它的
//		username := c.PostForm("username")
//		password := c.PostForm("password")
//		// 多选框
//		hobbys := c.PostFormArray("hobby")
//		c.String(http.StatusOK, fmt.Sprintf("type is %s,username is %s,password is %s,hobbys is %v",
//			type1, username, password, hobbys))
//	})
//	err := r.Run(":8100")
//	if err != nil {
//		return
//	}
//}

// upload
func main() {
	r := gin.Default()
	// url参数
	r.POST("/upload", func(c *gin.Context) {
		// 表单取文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		// 传到项目根目录,名字就用本身的
		c.SaveUploadedFile(file, file.Filename)
		// 打印信息
		c.String(200, fmt.Sprintf("'%s' upload!", file.Filename))
	})
	err := r.Run(":8100")
	if err != nil {
		return
	}
}
