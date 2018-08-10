package models

import (
	// "github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3" // _表示只进行初始化函数调用 其他函数暂不调用 因为是驱动，所以需要进行初始化调用(驱动的注册等操作)
	// "os"
	// "path"
	// "fmt"
	"strconv"
	"time"
)

const (
	_DB_NAME        = "bbs_test.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type GoBlog_Category struct {
	Id              int64
	Title           string
	Created_date    time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time
	TopicCount      int64
	TopicLastUserId int64
}

type GoBlog_Topic struct {
	Id              int64
	Uid             int64
	Author          string
	Title           string
	Content         string `orm:"size(65000)"`
	Attachment      string
	Created_date    time.Time `orm:"index"`
	Updated_date    time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
	ReplyLastTime   time.Time `orm:"index"`
}

func RegisterDb() {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:123456@/default?charset=utf8", 30)
	// 数据库文件不存在的时候，就创建
	// if !com.IsExist(_DB_NAME) {
	// 	// 递归的创建目录
	// 	os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
	// 	os.Create(_DB_NAME)
	// }
	orm.RegisterModel(new(GoBlog_Category), new(GoBlog_Topic))

	orm.RunSyncdb("default", false, true)
	// orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 3)
}

func MyRegister() {
	orm.Debug = true

	orm.RegisterDataBase("default", "mysql", "root:123456@/bbs_test?charset=utf8", 30)
	orm.RegisterModel(new(GoBlog_Category), new(GoBlog_Topic))
	orm.RunSyncdb("default", false, true)
}

func AddTopic(title, content string) error {
	o := orm.NewOrm()
	topic := &GoBlog_Topic{
		Title:        title,
		Content:      content,
		Created_date: time.Now(),
		Updated_date: time.Now(),
	}
	_, err := o.Insert(topic)
	return err
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &GoBlog_Category{Title: name}
	qs := o.QueryTable("go_blog__category")
	err := qs.Filter("title", name).One(cate)
	if err != nil {
		return err
	}
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

// 获取所有文章
func GetAllTopics(isDesc bool) ([]*GoBlog_Topic, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("go_blog__topic")
	topics := make([]*GoBlog_Topic, 0)
	if isDesc {
		qs = qs.OrderBy("-created_date")
	}
	_, err := qs.All(&topics)

	return topics, err
}

// 获取1篇文章
func GetOneTopic(tid string) (*GoBlog_Topic, error) {
	intTid, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	qs := o.QueryTable("go_blog__topic")
	topic := new(GoBlog_Topic)
	err = qs.Filter("id", intTid).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)

	return topic, err
}

// 更新一篇文章
func UpdateTopic(tid, title, content string) error {
	intTid, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &GoBlog_Topic{Id: intTid}
	if o.Read(topic) == nil {
		topic.Id = intTid
		topic.Title = title
		topic.Content = content
		topic.Updated_date = time.Now()
		intTid, err = o.Update(topic)
	}
	return err
}

// 删除一篇文章
func DeleteTopic(tid string) error {
	intTid, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &GoBlog_Topic{Id: intTid}
	_, err = o.Delete(topic)
	return err
}

// 获取所有分类
func GetAllCategories() ([]*GoBlog_Category, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("go_blog__category")
	cates := make([]*GoBlog_Category, 0)
	_, err := qs.All(&cates)

	return cates, err
}

func DeleteCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &GoBlog_Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}
