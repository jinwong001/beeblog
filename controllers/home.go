package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	topics,err:=models.GetAllTopics(true)
	if err!=nil{
		beego.Error(err)
	}
	c.Data["Topics"] = topics
	c.TplName = "home.html"
}

