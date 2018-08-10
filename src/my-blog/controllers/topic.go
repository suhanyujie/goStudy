package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"my-blog/models"
)

type TopicController struct {
	beego.Controller
}

// @router /v1/article/list [get]
func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = CheckAcount(this.Ctx)
	topics, err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["Topics"] = topics
	this.TplName = "Topic/topic.html"
}
// @router /v1/article/add [post]
func (this *TopicController) Add() {
	this.TplName = "Topic/topic_add.html"
	// this.Ctx.WriteString("add~~")
}

// 文章详情页
func (this *TopicController) Show() {
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = CheckAcount(this.Ctx)
	tid := this.Ctx.Input.Params()["0"]
	// beego.Error(this.Ctx.Input.Params["0"])
	fmt.Println("123123--")
	topic, err := models.GetOneTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Tid"] = tid

	this.TplName = "Topic/topic_show.html"
}

// 发表文章
func (this *TopicController) Post() {
	if !CheckAcount(this.Ctx) {
		this.Redirect("login", 302)
		return
	}
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	// 没有这个键的值，则是空字串
	tid := this.Input().Get("tid")
	var err error
	if len(tid) > 0 {
		err = models.UpdateTopic(tid, title, content)
	} else {
		err = models.AddTopic(title, content)
	}
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
}

// 文章详情页
func (this *TopicController) Edit() {
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = CheckAcount(this.Ctx)
	tid := this.Ctx.Input.Params()["0"]
	topic, err := models.GetOneTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Tid"] = tid

	this.TplName = "Topic/topic_edit.html"
}

// 文章的删除
func (this *TopicController) Delete() {
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = CheckAcount(this.Ctx)
	tid := this.Ctx.Input.Params()["0"]
	err := models.DeleteTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Redirect("/topic", 302)
}
