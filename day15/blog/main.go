package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sealandsigh/gotest5/day15/blog/controller"
	"github.com/sealandsigh/gotest5/day15/blog/dao/db"
)

func main() {
	router := gin.Default()
	dns := "root:zhujiajun@tcp(localhost:3306)/goday10?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	// 加载静态文件
	router.Static("/static/", "./static")
	// 加载模板
	router.LoadHTMLGlob("views/*")
	router.GET("/", controller.IndexHandle)
	router.GET("/category", controller.CategoryList)
	_ = router.Run(":8100")
}
