package controllers

import(
	"lovehome/models"
	"github.com/astaxie/beego"
)

type SessionControllers struct{
	beego.Controller
}

//把封装好的json字符串返回给前端
func(this *SessionControllers)RetData(resp interface{}){
	this.Data["json"]=resp
	this.ServeJSON()
}

//用户退出业务
func (this *SessionControllers)DelSessionName(){
	beego.Info("DelSessionName...DelSessionName....")
	resp:=make(map[string]interface{})

	resp["erron"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

	//将session删除
	this.DelSession("name")
	this.DelSession("user_id")
	this.DelSession("mobile")
}

func (this *SessionControllers)GetSessionName(){
	beego.Info("GetSessionName......")

	resp:=make(map[string]interface{})
	resp["errno"]=models.RECODE_SESSIONERR
	resp["errmsg"]=models.RecodeText(models.RECODE_SESSIONERR)

	defer this.RetData(resp)

	name_map:=make(map[string]interface{})
	name:=this.GetSession("name")

	if name!=nil{
		resp["erron"]=models.RECODE_OK
		resp["errmsg"]=models.RecodeText(models.RECODE_OK)
		name_map["name"]=name.(string)
		resp["data"]=name_map
	}

	return
}
