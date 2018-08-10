package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 入口函数
func main() {
	db,err:=sql.Open("mysql", "root:123456@/node?charset=utf8")
	if err!=nil{
		panic(err)
	}
	stmt,err := db.Prepare("INSERT node_article SET title=?,content=?")
	if err!=nil{
		panic(err)
	}
	res,err:= stmt.Exec("标题051001","this is content")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", res)
}

