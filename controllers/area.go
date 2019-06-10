package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"iHome/models"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) RetData(resp *map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *AreaController) GetArea() {
	//从session拿数据

	//从查询数据库
	or := orm.NewOrm()

	areas := []models.Area{}
	resp := map[string]interface{}{}
	defer c.RetData(&resp)

	n, err := or.QueryTable("area").All(&areas)
	if err != nil {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)

		return
	}

	if n == 0 {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)

		return
	}

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = areas

	beego.Info(resp)
}
