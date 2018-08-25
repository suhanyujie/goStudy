package beegoOrm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"log"
	"lesson/gorm/beegoOrm/models"
	"github.com/astaxie/beego"
)

func init() {
	beego.LoadAppConfig("ini", "src/lesson/gorm/beegoOrm/config/database.conf")
	connectionStr := beego.AppConfig.String("connection1")
	orm.RegisterDataBase("default", "mysql", connectionStr, 30)
	orm.RegisterModel(new(models.NovelList), new(models.NovelContent), new(models.NovelMain))
	//orm.RunSyncdb("default", false, true)
}

//查询多条
func Test2() []orm.Params {
	ormObject := orm.NewOrm()
	var lists []orm.Params
	qs := ormObject.QueryTable("NovelList")
	_, err := qs.Filter("novel_id", 2).Limit(10, 0).Values(&lists)
	if err != nil {
		log.Fatal(err)
	}
	for _, list := range lists {
		fmt.Println(list["Title"])
	}
	return lists
}

//普通查询
func Test1() {
	ormObject := orm.NewOrm()
	var lists []orm.Params
	num, err := ormObject.Raw("select * from novel_list limit 2").Values(&lists)
	if err != nil {
		log.Fatal(err)
	}
	for _, term := range lists {
		fmt.Println(term["id"], ":", term["title"])
	}
	fmt.Println(num)
}
