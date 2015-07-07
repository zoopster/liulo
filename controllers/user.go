// Package controllers

package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

// UserController is operations about User
type UserController struct {
	beego.Controller
}

func (this *UserController) SignUp() {
	fmt.Printf("%+v\n", this.Ctx.Input.RequestBody)

	this.Data["json"] = map[string]string{
		"name": "UserController#SignUp",
	}

	this.ServeJson()
}

func (this *UserController) SignIn() {

	this.Data["json"] = map[string]string{
		"name": "UserController#SignIn",
	}
	this.ServeJson()
}

func (this *UserController) GetCurrentUser() {

	this.Data["json"] = map[string]string{
		"name": "UserController#Me",
	}
	this.ServeJson()
}
