package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // init()
)

// Go 连接MySQL示例

var db *sql.DB // 是一个连接池对象

func initDB() (err error) {
	// 数据库信息
	// 用户名:密码@tcp(ip:端口)/数据库名字
	dsn := "root:zhujiajun@tcp(127.0.0.1:3306)/sql_test"
	// 连接数据库
	db, err = sql.Open("mysql", dsn) //不会校验用户名密码是否正确
	if err != nil {                  // dsn格式不正确的时候报错
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	// 设置数据库连接池最大连接数
	db.SetMaxOpenConns(10)
	return
}

// type user struct {
// 	id   int
// 	name string
// 	age  int
// }

func transactionDemo() {
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed,err:%v\n", err)
		return
	}
	// 执行多个SQL操作
	sqlStr1 := `update user set age=age-2 where id=2`
	sqlStr2 := `update user set age=age+1 where id=3`
	// 执行sql语句1
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		//要回滚
		tx.Rollback()
		fmt.Println("执行sql1出错了，要回滚")
		return
	}
	// 执行sql语句2
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		//要回滚
		tx.Rollback()
		fmt.Println("执行sql2出错了，要回滚")
		return
	}
	// 上面两步sql都执行成功了，提交本次事务
	err = tx.Commit()
	if err != nil {
		//要回滚
		tx.Rollback()
		fmt.Println("提交出错了，要回滚")
		return
	}
	fmt.Println("事务执行成功")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v", err)
	}
	fmt.Println("连接数据库成功！")
	transactionDemo()
}
