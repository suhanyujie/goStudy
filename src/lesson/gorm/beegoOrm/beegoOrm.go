package beegoOrm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"log"
	"lesson/gorm/beegoOrm/models"
	"github.com/astaxie/beego"
	string2 "practice/http/libs/string"
)

func init() {
	beego.LoadAppConfig("ini", "src/lesson/gorm/beegoOrm/config/database.conf")
	//connectionStr := beego.AppConfig.String("connection1")
	err := orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/bbs_test?charset=utf8", 30)
	if err != nil {
		log.Fatal(err)
	}
	orm.RegisterModel(new(models.NovelList), new(models.NovelContent), new(models.NovelMain))
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatal(err)
	}
}

type FictionOneOfList struct {
	Title string
	Url   string
}

/**
插入列表数据
 */
func InsertList(OneTask FictionOneOfList)(int,string) {
	chanpterNum := string2.Chinese2Int(OneTask.Title)
	list := &models.NovelList{
		NovelId:11,
		Url:OneTask.Url,
		Title:OneTask.Title,
		Chapter:chanpterNum,
	}
	ormObject := orm.NewOrm()
	id,err := ormObject.Insert(list)
	if err!=nil {
		log.Fatal(err)
		return 0,"新增失败！"
	}

	return int(id),"新增成功！"
}

//查询多条
func Test21() []orm.Params {
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

func Test2() {
	//status,msg := InsertList()
	//fmt.Print(status,msg)
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
