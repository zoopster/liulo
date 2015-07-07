package controllers

import "github.com/astaxie/beego"

// EventController is based on Beego Controller
type EventController struct {
	beego.Controller
}

func (this *EventController) Post() {

	this.Data["json"] = map[string]string{
		"name": "EventController#Post",
	}

	this.ServeJson()
}

func (this *EventController) GetBySlug() {

	this.Data["json"] = map[string]string{
		"name": "EventController#GetBySlug",
	}

	this.ServeJson()
}

func (this *EventController) GetById() {

	this.Data["json"] = map[string]string{
		"name": "EventController#GetById",
	}
	this.ServeJson()
}
