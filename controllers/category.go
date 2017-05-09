package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
	}

	c.Data["IsLogin"]=checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	//	c.TplName = "category.html"
	var err error
	c.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "category.html"

}
