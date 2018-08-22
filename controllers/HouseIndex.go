package controllers

import(
	"lovehome/models"
	"github.com/astaxie/beego"
)

type HouseIndexControllers struct{
	beego.Controller
}

func (this *HouseIndexControllers)RetData(resp interface{}){
	this.Data["json"]=resp
	this.ServeJSON()
}

func (this *HouseIndexControllers)GetHouseIndex(){
	beego.Info("GetHouseIndex.......")
	resp:=make(map[string]interface{})
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)
	return
}
