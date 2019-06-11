package models

import (
	"github.com/astaxie/beego"
	"github.com/weilaihui/fdfs_client"
	_ "github.com/weilaihui/fdfs_client"
)

func UploadByFilename(fileName string) (groupName, fileId string, err error) {
	fdfsClient, errClient := fdfs_client.NewFdfsClient("./conf/client.conf")
	if errClient != nil {
		beego.Info("New FdfsClient error %s", errClient.Error())
		return "", "", errClient
	}

	uploadResponse, errUpload := fdfsClient.UploadByFilename(fileName)
	if errUpload != nil {
		beego.Info("UploadByfilename error %s", errUpload.Error())
		return "", "", errUpload
	}
	beego.Info(uploadResponse.GroupName)
	beego.Info(uploadResponse.RemoteFileId)

	//删除文件
	//fdfsClient.DeleteFile(uploadResponse.RemoteFileId)

	return uploadResponse.GroupName, uploadResponse.RemoteFileId, nil
}
