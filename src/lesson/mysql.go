package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	db, err := sql.Open("mysql", "root:123456@/node")
	// id := 2;
	result, err := db.Query("select title,content from node_article")
	if err != nil {
		log.Fatal(err)
	}
	var title, content string
	for result.Next() {
		if err := result.Scan(&title, &content); err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(title), content)
	}
	defer db.Close()
}
