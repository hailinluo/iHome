package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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

func (this *SessionController) Login() {
	//得到用户信息
	req := make(map[string]interface{})
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err != nil {
		beego.Info(err.Error())
		return
	}
	beego.Info(req)  //[mobile:444 password:444

	resp := map[string]interface{}{}
	defer this.RetData(&resp)

	//判断是否合法
	if req["mobile"] == nil || req["password"] == nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	//查询数据库，判断账号密码是否存在
	o := orm.NewOrm()
	user := models.User{}
	query := o.QueryTable("user")
	if err := query.Filter("mobile", req["mobile"]).One(&user); err != nil {
	//if err := o.Read(&user); err != nil {
		resp["errno"] = models.RECODE_LOGINERR
		resp["errmsg"] = models.RecodeText(models.RECODE_LOGINERR)
		return
	}
	if user.Password_hash != req["password"] {
		resp["errno"] = models.RECODE_PARAMERR
		resp["errmsg"] = models.RecodeText(models.RECODE_PARAMERR)
		return
	}

	//添加session
	this.SetSession("name", user.Name)
	this.SetSession("mobile", user.Mobile)
	this.SetSession("user_id", user.Id)

	//返回json
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}
