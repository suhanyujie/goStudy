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
	InsertId int
	Chapter int
}

//查询列表
func QueryList(params map[string]string) (models.NovelList,error) {
	o := orm.NewOrm()
	qs := o.QueryTable("novel_list")
	if novel_id,ok:=params["novel_id"];ok {
		qs = qs.Filter("novel_id", novel_id)
	}
	if chapter,ok:=params["chapter"];ok {
		qs = qs.Filter("chapter", chapter)
	}
	var oneList models.NovelList
	err := qs.One(&oneList)
	if err!=nil {
		log.Println(err)
	}
	return oneList,nil
}

/**
插入列表数据
 */
func InsertList(OneTask FictionOneOfList)(int,string) {
	list := &models.NovelList{
		NovelId:12,
		Url:OneTask.Url,
		Title:OneTask.Title,
		Chapter:OneTask.Chapter,
	}
	ormObject := orm.NewOrm()
	id,err := ormObject.Insert(list)
	if err!=nil {
		log.Fatal(err)
		return 0,"新增失败！"
	}

	return int(id),"新增成功！"
}

/**
插入详细内容数据
 */
func InsertDetail(detail models.NovelContent)(int,string) {
	ormObject := orm.NewOrm()
	id,err := ormObject.Insert(&detail)
	if err!=nil {
		log.Fatal(err)
		return 0,"新增详细内容失败！"
	}

	return int(id),"新增详细内容成功！"
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
