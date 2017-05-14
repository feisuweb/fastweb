package main

import (
	"github.com/astaxie/beego"
	"github.com/feisuweb/fastweb/models"
	_ "github.com/feisuweb/fastweb/routers"
	"os"
)

func main() {
	//创建附件目录
	os.Mkdir("upload", os.ModePerm)
	os.Mkdir("html", os.ModePerm)
	os.Mkdir("upload/images", os.ModePerm)
	os.Mkdir("upload/files", os.ModePerm)
	beego.AddFuncMap("ReplaceMobile", models.ReplaceMobile)
	beego.AddFuncMap("GetMemberTypeName", models.GetMemberTypeNameById)
	beego.AddFuncMap("GetMemberTypeName", models.GetMemberTypeNameById)
	beego.AddFuncMap("GetMemberNameById", models.GetMemberNameById)
	beego.SetLogFuncCall(true)
	beego.SetLogger("file", `{"filename":"logs/web.log"}`)
	beego.Info("服务已经启动...")
	beego.Run()
}
