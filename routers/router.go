package routers

import (
	"lovehome/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	//请求地域
	beego.Router("/api/v1.0/areas",&controllers.AreaControllers{},"get:GetAreaInfo")
	//请求session
	beego.Router("/api/v1.0/session",&controllers.SessionControllers{},"get:GetSessionName")
	//请求index
	beego.Router("/api/v1.0/houses/index",&controllers.HouseIndexControllers{},"get:GetHouseIndex")
	//请求注册
	beego.Router("/api/v1.0/users",&controllers.UserControllers{},"post:Reg")
	//请求登入
	beego.Router("/api/v1.0/sessions",&controllers.UserControllers{},"post:Login")
	//用户上传头像业务
	//beego.Router("/api/v1.0/user/avatar",&controllers.UserControllers{},"post:UploadAvatar")
}
