package controllers

import (
	"github.com/riveryang/sysgo/models"
)

type SysController struct {
	BaseController
}

// @Title Get
// @Description find pc stat
// @router / [get]
func (o *SysController) Get() {
	o.AllowCross()
	if pcStat, err := models.NewPcStat(); err != nil {
		o.Data["json"] = err.Error()
	} else {
		if encrypt, err := pcStat.Encrypt(); err != nil {
			o.Data["json"] = err.Error()
		} else {
			var obj = make(map[string]string)
			obj["stat"] = encrypt
			o.Data["json"] = obj
		}
	}

	o.ServeJSON()
}
