package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"iHome/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) RetData(resp *map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *UserController) Register() {
	req := make(map[string]interface{})
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err != nil {
		beego.Info(err.Error())
		return
	}

	//{mobile: "111", password: "444", sms_code: "333"
	or := orm.NewOrm()
	user := models.User{}
	user.Password_hash = req["password"].(string)
	user.Name = req["mobile"].(string)
	user.Mobile = req["mobile"].(string)

	resp := map[string]interface{}{}
	defer this.RetData(&resp)

	_, err := or.Insert(&user)
	if err != nil {
		resp["errno"] = models.RECODE_USERERR
		resp["errmsg"] = models.RecodeText(models.RECODE_USERERR)

		return
	}

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	this.SetSession("name", user.Name)
}
