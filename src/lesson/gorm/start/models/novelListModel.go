package models

import (
	"github.com/jinzhu/gorm"
	"log"
	"fmt"
)

type NovelList struct {
	gorm.Model
	Id       uint   `gorm:"primary_key"`
	Novel_id int    `gorm:"size:11"`
	Url      string `gorm:"size:50"`
	Chapter  int
	Title    string `gorm:"size:50"`
	Flag     int
	Err_flag int
}

// todo 获取表名称
func (_this NovelList) TableName() string {
	return "novel_list"
}


func Test1() {
	db, err := gorm.Open("mysql", "root:123456@/bbs_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println(db.HasTable("novel_list"))

}
