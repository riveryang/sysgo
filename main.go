package main

import (
	"github.com/astaxie/beego"
	_ "github.com/riveryang/sysgo/routers"
)

func init () {
	beego.BConfig.AppName = "sysgo"
	beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	beego.BConfig.Listen.HTTPPort = 35791
	beego.BConfig.RunMode = "prod"
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.EnableDocs = true
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
