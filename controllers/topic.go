package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
)

type TopicController  struct {
	beego.Controller
}

func (c *TopicController)Get() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 301)
		return
	}
	c.Data["IsLogin"] = true
	c.Data["IsTopic"] = true
	var err error
	c.Data["Topics"], err = models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "topic.html"

}

func (c *TopicController)Post() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 301)
		return
	}

	tid := c.Input().Get("tid")
	title := c.Input().Get("title")
	content := c.Input().Get("content")

	beego.Debug("tid:     " + tid)

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, content)
	} else {
		err = models.ModifyTopic(tid, title, content)
	}
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/topic", 301)

}

func (c *TopicController)Add() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 301)
		return
	}
	c.Data["IsLogin"] = true
	c.Data["IsTopic"] = true
	c.TplName = "topic_add.html"
}

func (c *TopicController)Delete() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 301)
		return
	}
	c.Data["IsLogin"] = true
	c.Data["IsTopic"] = true
	err := models.DelTopic(c.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic", 301)
}

func (c *TopicController)Modify() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 301)
		return
	}
	tid := c.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 301)
		return
	}
	c.Data["Topic"] = topic
	c.Data["IsLogin"] = true
	c.Data["IsTopic"] = true
	c.Data["Tid"] = tid
	c.TplName = "topic_modify.html"
}

func (c *TopicController)View() () {
	if tid, ok := c.Ctx.Input.Params()["0"]; ok {
		topic, err := models.GetTopic(tid)
		if err != nil {
			beego.Error(err)
			c.Redirect("/", 301)
			return
		}
		c.Data["Topic"] = topic
	} else {
		c.Redirect("/", 301)
		return
	}
	c.TplName = "topic_view.html"

}