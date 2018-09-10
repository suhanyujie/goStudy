package models

import "github.com/astaxie/beego/orm"

type NovelMain struct {
	Id          uint
	ListUrl     string
	BaseUrl     string
	NovelStatus uint
	Desc        string
	InsertDate  orm.DateField `orm:"auto_now"`
}

type NovelList struct {
	Id       uint
	NovelId  uint
	Url      string
	Chapter  int
	Title    string
	Flag     int
	Err_flag int
}

type NovelContent struct {
	Id       uint
	ListId   uint
	Chapter  int
	Title    string
	Content  string
	WorkerId uint
	Date     orm.DateField `orm:"auto_now"`
	ErrFlag  uint
}
