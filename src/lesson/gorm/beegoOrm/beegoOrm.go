package beegoOrm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"log"
)


type NovelList struct {
	Id       uint   `gorm:"primary_key"`
	Novel_id int    `gorm:"size:11"`
	Url      string `gorm:"size:50"`
	Chapter  int
	Title    string `gorm:"size:50"`
	Flag     int
	Err_flag int
}

func init() {
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/bbs_test?charset=utf8", 30)
	orm.RegisterModel(new(NovelList))
	orm.RunSyncdb("default", false, true)
}

func Test1() {
	o1 := orm.NewOrm()
	var lists []orm.Params
	num,err := o1.Raw("select * from novel_list limit 2").Values(&lists)
	if err!=nil {
		log.Fatal(err)
	}
	for _,term := range lists {
		fmt.Println(term["id"],":",term["title"])
	}
	fmt.Println(num)
}
