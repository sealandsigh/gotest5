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

type user struct {
	id   int
	name string
	age  int
}

// 查询单条记录
func queryOne(id int) {
	var u1 user
	//1 写查询单条记录的sql语句
	sqlStr := `select id,name,age from user where id=?;`
	//2 执行
	// rowObj := db.QueryRow(sqlStr, 1) // 从连接池里面拿一个连接去数据库查询单条语句
	//3 拿到结果
	// rowObj.Scan(&u1.id, &u1.name, &u1.age) // 必须对rowObj对象调用Scan方法，因为该方法会释放数据库连接
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)
	// 打印结果
	fmt.Printf("u1:%#v\n", u1)
}

func insert() {

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v", err)
	}
	fmt.Println("连接数据库成功！")
	queryOne(2)
}
