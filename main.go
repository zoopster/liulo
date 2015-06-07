package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	_ "github.com/dwarvesf/liulo/routers"
)

const (
	APP_VER = "0.1"
)

func main() {

	beego.EnableAdmin = true

	// Configure env
	beego.Info(beego.AppName, APP_VER)

	// Register template functions.
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}
