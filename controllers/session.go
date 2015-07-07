package controllers

import "github.com/astaxie/beego"

type SessionController struct {
	beego.Controller
}

func (this *SessionController) Post() {
	this.Data["json"] = map[string]string{
		"name": "SessionController#Post",
	}
	this.ServeJson()
}

func (this *SessionController) GetById() {
	this.Data["json"] = map[string]string{
		"name": "SessionController#GetById",
	}
	this.ServeJson()
}

func (this *SessionController) GetAll() {
	this.Data["json"] = map[string]string{
		"name": "SessionController#GetAll",
	}
	this.ServeJson()
}
