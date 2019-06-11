package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"iHome/models"
	"time"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) RetData(resp *map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *AreaController) GetArea() {
	areas := []models.Area{}
	resp := map[string]interface{}{}
	defer c.RetData(&resp)

	//从Redis缓存中拿数据
	redisConn, _ := cache.NewCache("redis", `{"key":"ihome", "conn":"127.0.0.1:6379", "dbNum":"0"}`)
	if redisConn != nil {
		if areasData := redisConn.Get("areas"); areasData != nil {
			resp["data"] = areasData
			return
		}
	}

	//从查询数据库
	or := orm.NewOrm()

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
	//beego.Info(resp)

	//把数据转换成json存入缓存
	if redisConn != nil {
		jsonBytes, err := json.Marshal(areas)
		if err != nil {
			resp["errno"] = models.RECODE_DATAERR
			resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
			return
		}

		errPut := redisConn.Put("areas", string(jsonBytes), time.Second * 3600)
		if errPut != nil {
			beego.Error("redis put err.")
		}
	}


}
