package controllers

import (
	"github.com/astaxie/beego"
	"iHome/models"
)

type SessionController struct {
	beego.Controller
}

func (this *SessionController) RetData(resp *map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *SessionController) GetSessionData() {
	resp := map[string]interface{}{}
	defer c.RetData(&resp)

	resp["errno"] = models.RECODE_SESSIONERR
	resp["errmsg"] = models.RecodeText(models.RECODE_SESSIONERR)

	user := models.User{}
	name := c.GetSession("name")
	if name != nil {
		user.Name = name.(string)
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		resp["data"] = user
	}
}

func (this *SessionController) DelSessionData() {
	resp := map[string]interface{}{}
	defer this.RetData(&resp)

	this.DelSession("name")

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}
