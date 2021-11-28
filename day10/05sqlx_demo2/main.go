package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // init()
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB // 是一个连接池对象

func initDB() (err error) {
	// 数据库信息
	// 用户名:密码@tcp(ip:端口)/数据库名字
	dsn := "root:zhujiajun@tcp(127.0.0.1:3306)/sql_test"
	// 连接数据库
	db, err = sqlx.Open("mysql", dsn) //不会校验用户名密码是否正确
	if err != nil {                   // dsn格式不正确的时候报错
		return
	}
	// 设置数据库连接池最大连接数
	db.SetMaxOpenConns(10)
	return
}

type user struct {
	ID   int
	Name string
	Age  int
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	sqlStr1 := `select id,name,age from user where id=3`
	var u user
	db.Get(&u, sqlStr1)
	fmt.Printf("u:%v\n", u)
}
