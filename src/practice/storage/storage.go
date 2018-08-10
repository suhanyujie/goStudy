package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

type StorageObj struct {
	db *(sql.DB)//数据库操纵对象
	lastSql string//上一个执行的sql语句
}

type GoSpiderListModel struct {
	Id string
	Website string
	FictionId int64
	Title string
	Url string
	Status int8
	AddTime string
	UpdateTime string
}


// todo
func (_this *StorageObj) Connect()  {
	db,err := sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/novel_spider?charset=utf8")
	if err!= nil {
		log.Fatal(err)
	}
	_this.db = db
}

//新增数据
func (_this *StorageObj)Insert(sqlString string) (id int64) {
	db := _this.db
	res,err := db.Exec(sqlString)
	_this.lastSql = sqlString
	if err!= nil {
		log.Fatal(err)
	}
	id,err = res.LastInsertId()
	if err!= nil {
		log.Fatal(err)
	}
	defer db.Close()
	return id
}

//根据条件查询数据的一个字段值
func (_this *StorageObj) GetOne(sqlString string) string  {
	Id := ""
	_this.db.QueryRow(sqlString).Scan(&Id)
	fmt.Println("value is :"+Id);
	return Id;
}

// todo
func (_this *StorageObj) GetRow(sqlString string) (*GoSpiderListModel) {
	data := new(GoSpiderListModel)
	err := _this.db.QueryRow(sqlString).Scan(&data.Id,&data.FictionId,&data.Url,&data.AddTime,&data.Status,&data.Website,&data.UpdateTime,&data.Title)
	if err!=nil {
		log.Fatal(err)
	}

	return data
}

//更新数据

//根据条件查询数据是否存在

// todo 获取最后一次执行的sql语句
func (_this *StorageObj) GetLastSql() (sql string){
	return _this.lastSql
}


