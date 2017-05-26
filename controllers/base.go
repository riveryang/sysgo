package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (c *BaseController) AllowCross() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("content-type", "application/json")
}