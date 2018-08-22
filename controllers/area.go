package controllers

import(
	"lovehome/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AreaControllers struct{
	beego.Controller
}

func (this *AreaControllers) RetData(resp interface{}){
	this.Data["json"]=resp
	this.ServeJSON()
}

func (this *AreaControllers) GetAreaInfo(){
	beego.Info("GetAreaInfo succ.....")
	resp:=make(map[string]interface{})
	resp["errno"]=0
	resp["errmsg"]="ok"

	defer this.RetData(resp)
	
	//思路
	//1.从缓冲区读取数据
	//2.如果redis中有之前的数据，直接返回给前端
	//3.如果redis中没有数据，从MySQL中查询
	o:=orm.NewOrm()

	var areas []models.Area
	qs:=o.QueryTable("area")
	num,err:=qs.All(&areas)
	if err!=nil{
		resp["errno"]=4001
		resp["errmsg"]="查询失败"
		return
	}

	if num==0{
		resp["errno"]=4002
		resp["errmsg"]="没有数据"
	}

	resp["data"]=areas
	return
}
