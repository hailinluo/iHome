package controllers

import (
	"github.com/astaxie/beego"
	"iHome/models"
)

type HouseIndexController struct {
	beego.Controller
}

func (this *HouseIndexController) RetData(resp *map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *HouseIndexController) GetHouseIndex() {
	resp := map[string]interface{}{}
	defer c.RetData(&resp)

	resp["errno"] = models.RECODE_DATAERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
}
