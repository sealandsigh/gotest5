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

//SQL注入

//sql 注入示例
func sqlInjectDemo(name string) {
	// 自己拼接sql语句的字符串
	sqlStr := fmt.Sprintf("select id,name,age from user where name='%s'", name)

	fmt.Printf("SQL:%s\n", sqlStr)

	var users []user
	err := db.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	for _, u := range users {
		fmt.Printf("user:%#v\n", u)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	// sql注入的几种示例
	// sqlInjectDemo("朱嘉骏")
	// 这个注入相当于可以查到数据库中所有的数据，就是俗称的拖库
	// sqlInjectDemo("xxx' or 1=1 #")
	sqlInjectDemo("xxx' union select * from user #")

	// 解决sql注入最关键的就是不要自己拼接字符串
	// 利用预编译
}
