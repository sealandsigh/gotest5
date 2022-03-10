package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Login struct {
	// binding: "required" 修饰字段，若接收为空值，则报错，是必须字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

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

//// upload 单文件提交
//func main() {
//	r := gin.Default()
//	// url参数
//	// 限制表达上传大小 8MB, 默认为32MB
//	r.MaxMultipartMemory = 8 << 20
//	r.POST("/upload", func(c *gin.Context) {
//		// 表单取文件
//		file, _ := c.FormFile("file")
//		log.Println(file.Filename)
//		// 传到项目根目录,名字就用本身的
//		c.SaveUploadedFile(file, file.Filename)
//		// 打印信息
//		c.String(200, fmt.Sprintf("'%s' upload!", file.Filename))
//	})
//	err := r.Run(":8100")
//	if err != nil {
//		return
//	}
//}

//// upload 多文件提交
//func main() {
//	r := gin.Default()
//	// url参数
//	// 限制表达上传大小 8MB, 默认为32MB
//	r.MaxMultipartMemory = 8 << 20
//	r.POST("/upload", func(c *gin.Context) {
//		form, err := c.MultipartForm()
//		if err != nil {
//			c.String(http.StatusBadGateway, fmt.Sprintf("get err %s", err.Error()))
//		}
//		// 获取所有图片
//		files := form.File["files"]
//		// 遍历所有图片
//		for _, file := range files {
//			// 逐个存
//			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
//				c.String(http.StatusBadRequest, fmt.Sprintf("upload error %s", err.Error()))
//				return
//			}
//		}
//		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
//	})
//	err := r.Run(":8100")
//	if err != nil {
//		return
//	}
//}

//// 路由组
//func main() {
//	r := gin.Default()
//	// 路由组1, 处理GET请求
//	v1 := r.Group("/v1")
//	// {} 这个括号是书写规范
//	{
//		v1.GET("/login", login)
//		v1.GET("/submit", submit)
//	}
//	v2 := r.Group("v2")
//	{
//		v2.POST("/login", login)
//		v2.POST("/submit", submit)
//	}
//	err := r.Run(":8100")
//	if err != nil {
//		return
//	}
//}
//
//func login(c *gin.Context) {
//	name := c.DefaultQuery("name", "jack")
//	c.String(200, fmt.Sprintf("hello %s\n", name))
//}
//
//func submit(c *gin.Context) {
//	name := c.DefaultQuery("name", "jack")
//	c.String(200, fmt.Sprintf("hello %s\n", name))
//}

// 路由原理
// httprouter 会将所有路由规则构造一棵前缀树
// 例如有 root and as at cn com
// 树的格式便于查询

//// 数据绑定,以下是json数据绑定方式
//func main() {
//	// 1. 创建路由
//	// 默认使用了2个中间件Logger(), Recover()
//	r := gin.Default()
//	// json绑定
//	r.POST("loginJSON", func(c *gin.Context) {
//		// 声明接收的变量
//		var json Login
//		// 将request的body中的数据，自动按照json格式解析到结构体
//		if err := c.ShouldBindJSON(&json); err != nil {
//			// 返回错误信息
//			// gin.H 封装了生成json数据的工具
//			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
//			return
//		}
//		// 判断用户名密码是否正确
//		if json.User != "root" || json.Password != "admin" {
//			c.JSON(http.StatusBadGateway, gin.H{"status": "304"})
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{"status": "200"})
//	})
//	err := r.Run(":8100")
//	if err != nil {
//		return
//	}
//}

//// 数据绑定,以下是form数据绑定方式
//func main() {
//	// 1. 创建路由
//	// 默认使用了2个中间件Logger(), Recover()
//	r := gin.Default()
//	// json绑定
//	r.POST("loginForm", func(c *gin.Context) {
//		// 声明接收的变量
//		var form Login
//		// Bind()默认解析并绑定form格式
//		// 根据请求头中的content-type自动推断
//		if err := c.Bind(&form); err != nil {
//			// 返回错误信息
//			// gin.H 封装了生成json数据的工具
//			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
//			return
//		}
//		// 判断用户名密码是否正确
//		if form.User != "root" || form.Password != "admin" {
//			c.JSON(http.StatusBadGateway, gin.H{"status": "304"})
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{"status": "200"})
//	})
//	err := r.Run(":8100")
//	if err != nil {
//		return
//	}
//}

//// 数据绑定,以下是uri数据绑定方式
//func main() {
//	// 1. 创建路由
//	// 默认使用了2个中间件Logger(), Recover()
//	r := gin.Default()
//	// json绑定
//	r.GET("/:user/:password", func(c *gin.Context) {
//		// 声明接收的变量
//		var login Login
//		// Bind()默认解析并绑定form格式
//		// 根据请求头中的content-type自动推断
//		if err := c.ShouldBindUri(&login); err != nil {
//			// 返回错误信息
//			// gin.H 封装了生成json数据的工具
//			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
//			return
//		}
//		// 判断用户名密码是否正确
//		if login.User != "root" || login.Password != "admin" {
//			c.JSON(http.StatusBadGateway, gin.H{"status": "304"})
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{"status": "200"})
//	})
//	err := r.Run(":8100")
//	if err != nil {
//		return
//	}
//}

//// 数据响应,数据渲染
//func main() {
//	// 1. 创建路由
//	// 默认使用了2个中间件Logger(), Recover()
//	r := gin.Default()
//	// 1. json
//	r.GET("/someJSON", func(c *gin.Context) {
//		c.JSON(200, gin.H{"message": "someJSON", "status": "200"})
//	})
//	// 2. 结构体响应
//	r.GET("/someStruct", func(c *gin.Context) {
//		var msg struct {
//			Name    string
//			Message string
//			Number  int
//		}
//		msg.Name = "root"
//		msg.Message = "message"
//		msg.Number = 123
//		c.JSON(200, msg)
//	})
//	// 3. XML
//	r.GET("/someXML", func(c *gin.Context) {
//		c.XML(200, gin.H{"message": "abc"})
//	})
//	// 4. YAML响应
//	r.GET("/someYAML", func(c *gin.Context) {
//		c.YAML(200, gin.H{"name": "zhangsan"})
//	})
//	// 5. protobuf格式 谷歌开发的高效存储读取的工具
//	// 数组？切片？如果自己构建一个传输格式，应该是什么格式?
//	r.GET("/someProtoBuf", func(c *gin.Context) {
//		reps := []int64{int64(1), int64(2)}
//		// 定义数据
//		label := "label"
//		// 传protobuf格式数据
//		data := &protoexample.Test{
//			Label: &label,
//			Reps:  reps,
//		}
//		c.ProtoBuf(200, data)
//	})
//	err := r.Run(":8100")
//	if err != nil {
//		return
//	}
//}

//// html模板渲染
//// gin 支持加载html模板，然后根据模板参数进行配置并返回相应的数据，本质上就是字符串的转换
//// loadHTMLGlob()方法可以加载模板文件
//
//func main() {
//	// 1. 创建路由
//	// 默认使用了2个中间件Logger(), Recover()
//	r := gin.Default()
//	// 加载模板文件
//	r.LoadHTMLGlob("templates/*")
//	//r.LoadHTMLFiles("templates/index.tmpl")
//	r.GET("index", func(c *gin.Context) {
//		// 根据文件名渲染
//		// 最终json将title替换
//		c.HTML(200, "index.tmpl", gin.H{"title": "我的标题"})
//	})
//	err := r.Run(":8100")
//	if err != nil {
//		return
//	}
//}

//// 重定向
//func main() {
//	// 1. 创建路由
//	// 默认使用了2个中间件Logger(), Recover()
//	r := gin.Default()
//
//	r.GET("/redirect", func(c *gin.Context) {
//		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
//	})
//	err := r.Run(":8100")
//	if err != nil {
//		return
//	}
//}

// 异步和同步
// goroutine 机制可以方便的实现异步处理
// 另外，在启动新的groutine时，不应该使用原始的 上下文，必须使用它的只读副本
func main() {
	// 1. 创建路由
	// 默认使用了2个中间件Logger(), Recover()
	r := gin.Default()
	// 1. 异步
	r.GET("/long_async", func(c *gin.Context) {
		// 需要搞一个副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行" + copyContext.Request.URL.Path)
		}()
	})
	// 2. 同步
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行" + c.Request.URL.Path)
	})

	err := r.Run(":8100")
	if err != nil {
		return
	}
}
