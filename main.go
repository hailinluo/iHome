package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "iHome/models"
	_ "iHome/routers"
	"net/http"
	"strings"
)

func main() {
	//beego.Run("80")  指定端口
	ignoreStaticPath()
	//models.UploadByFilename("./README.md")
	beego.Run()
}

func ignoreStaticPath() {
	beego.SetStaticPath("group1/M00/", "fdfs/storage_data/data/")

	//透明static
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url: ", orpath)
	//如果请求URL带有api字段，不需要静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)
}
