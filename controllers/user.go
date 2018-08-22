package controllers

import(
	"lovehome/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

type UserControllers struct{
	beego.Controller
}

//把封装好的json字符串返回给前端
func (this *UserControllers)RetData(resp interface{}){
	this.Data["json"]=resp
	this.ServeJSON()
}

func (this *UserControllers)Reg(){
	beego.Info("Reg..........")

	resp:=make(map[string]interface{})
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

	//存储前端返回的json数据信息
	var regRequesMap=make(map[string]interface{})
	//1.得到客户端请求的json数据
	json.Unmarshal(this.Ctx.Input.RequestBody,&regRequesMap)
	beego.Info("用户名：",regRequesMap["mobile"])
	beego.Info("密码：",regRequesMap["password"])
	beego.Info("验证码：",regRequesMap["sms_code"])

	//2.判断数据是否合法
	if regRequesMap["mobile"]==""||regRequesMap["password"]==""||regRequesMap["sms_code"]==""{
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]=models.RecodeText(models.RECODE_REQERR)
		return
	}

	//3.将数据存入数据库
	user:=models.User{}
	//把从前端获取的手机号复制给结构体里的成员变量
	user.Mobile=regRequesMap["mobile"].(string)
	user.Password_hash=regRequesMap["password"].(string)
	user.Name=regRequesMap["mobile"].(string)

	//操作数据库
	o:=orm.NewOrm()
	id,err:=o.Insert(&user)
	if err!=nil{
		beego.Info("Insert mysql err")
		resp["erron"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		return
	}

	beego.Info("reg succ----------------------id=",id)
	//4.将当前用户的信息存储到session
	this.SetSession("name",user.Mobile)
	this.SetSession("user_id",id)
	this.SetSession("mobile",user.Mobile)

	return

}

func (this *UserControllers) Login(){
	beego.Info("Login succ----------------------")

	resp:=make(map[string]interface{})

	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)
	//存储前端返回的json数据信息
	var LoginRequesMap=make(map[string]interface{})

	//1.得到前端返回的json数据
	json.Unmarshal(this.Ctx.Input.RequestBody,&LoginRequesMap)
	beego.Info("用户名：",LoginRequesMap["mobile"])
	beego.Info("密码：",LoginRequesMap["password"])

	//2.判断数据是否合法
	if LoginRequesMap["mobile"]==""||LoginRequesMap["password"]==""{
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]=models.RecodeText(models.RECODE_REQERR)
		return
	}

	//3.查询数据库得到user
	var user models.User
	
	//操作数据库user
	o:=orm.NewOrm()
	qs:=o.QueryTable("user")

	if err:=qs.Filter("mobile",LoginRequesMap["mobile"]).One(&user);err!=nil{
		beego.Info("查询失败")
		resp["erron"]=models.RECODE_NODATA
		resp["errmsg"]=models.RecodeText(models.RECODE_NODATA)

		return
	}

	//4.对比密码
	if user.Password_hash!=LoginRequesMap["password"].(string){
		beego.Info("密码不对")
		resp["erron"]=models.RECODE_NODATA
		resp["errmsg"]=models.RecodeText(models.RECODE_NODATA)
		return
	}

	beego.Info("=========login succ ===user.name=",user.Name)

	//将当前用户信息存储到session
	this.SetSession("name",user.Mobile)
	this.SetSession("mobile",user.Mobile)
	this.SetSession("user_id",user.Id)
	return
}
