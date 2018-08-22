package main

import (
	_ "lovehome/routers"
	_ "lovehome/models"
	"net/http"
	"strings"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
)

func TransparentStatic(ctx *context.Context){
	orpath:=ctx.Request.URL.Path
	beego.Debug("request url: ",orpath)

	if strings.Index(orpath,"api")>=0{
		return
	}
	http.ServeFile(ctx.ResponseWriter,ctx.Request,"static/html/"+ctx.Request.URL.Path)
}

func ignoreStaticPath(){
	beego.InsertFilter("/",beego.BeforeRouter,TransparentStatic)
	beego.InsertFilter("/*",beego.BeforeRouter,TransparentStatic)
}
func main() {
	ignoreStaticPath()
	
	beego.SetStaticPath("/group1/M00","fastdfs/storage_data/data")

	beego.Run()
}

