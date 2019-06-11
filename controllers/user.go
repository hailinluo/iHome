package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/weilaihui/fdfs_client"
	"iHome/models"
	"path"
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

func (this *UserController) PostAvatar() {
	resp := map[string]interface{}{}
	defer this.RetData(&resp)

	fileData, fileHeader, errFile := this.GetFile("avatar")
	if errFile != nil {
		beego.Info("GetFile error %s", errFile.Error())
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	//获取文件后缀
	fileSuffix := path.Ext(fileHeader.Filename)

	//将文件存储到FastDFS
	fdfsClient, errClient := fdfs_client.NewFdfsClient("conf/client.conf")
	if errClient != nil {
		beego.Info("New FdfsClient error %s", errClient.Error())
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	fileBuffer := make([]byte, fileHeader.Size)
	_, errRead := fileData.Read(fileBuffer)
	if errRead != nil {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	uploadResponse, errUpload := fdfsClient.UploadByBuffer(fileBuffer, fileSuffix[1:])
	if errUpload != nil {
		beego.Info("TestUploadByBuffer error %s", errUpload.Error())
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	fmt.Println(uploadResponse.GroupName)
	fmt.Println(uploadResponse.RemoteFileId)
	//fdfsClient.DeleteFile(uploadResponse.RemoteFileId)

	//从session中拿到user_id
	userId := this.GetSession("user_id")

	//根据userId，将avatar链接插入User数据表中
	user := models.User{}
	o := orm.NewOrm()
	if o == nil {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	query := o.QueryTable("user")
	if query == nil {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	query.Filter("Id", userId).One(&user)
	user.Avatar_url = uploadResponse.RemoteFileId

	_, errUpdate := o.Update(&user)
	if errUpdate != nil {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	avatarUrlMap := make(map[string]string)
	avatarUrlMap["avatar_url"] = "http://192.168.110.17:8888/" + uploadResponse.RemoteFileId
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = avatarUrlMap
}
