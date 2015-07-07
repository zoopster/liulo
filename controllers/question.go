package controllers

import "github.com/astaxie/beego"

type QuestionController struct {
	beego.Controller
}

func (this *QuestionController) Post() {
	this.Data["json"] = map[string]string{
		"name": "QuestionController#Post",
	}

	this.ServeJson()
}

func (this *QuestionController) GetAll() {
	this.Data["json"] = map[string]string{
		"name": "QuestionController#GetAll",
	}
	this.ServeJson()
}

func (this *QuestionController) GetById() {
	this.Data["json"] = map[string]string{
		"name": "QuestionController#GetById",
	}
	this.ServeJson()
}
