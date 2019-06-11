package routers

import (
	"github.com/astaxie/beego"
	"iHome/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1.0/areas", &controllers.AreaController{}, "get:GetArea")
	beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{}, "get:GetHouseIndex")
	beego.Router("/api/v1.0/session", &controllers.SessionController{}, "get:GetSessionData;delete:DelSessionData")
	beego.Router("/api/v1.0/sessions", &controllers.SessionController{}, "post:Login")
	beego.Router("/api/v1.0/users", &controllers.UserController{}, "post:Register")

	//beego.Router("/api/v1.0/user", &controllers.UserController{}, "get:GetUser")
	beego.Router("/api/v1.0/user/avatar", &controllers.UserController{}, "post:PostAvatar")
}
