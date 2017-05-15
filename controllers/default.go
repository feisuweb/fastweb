package controllers

type MainController struct {
	baseController
}

func (this *MainController) Get() {
	if IsLogin {
		this.Data["PageTitle"] = "控制面板"
		this.Layout = "layout/_member_layout.html"
		this.TplName = "controller.html"
	} else {
		this.Data["PageTitle"] = "首页"
		this.TplName = "index.html"

	}
	this.Data["SiteName"] = "微信采集服务"
	this.Data["Keywords"] = "astaxie@gmail.com"
	this.Data["Description"] = "astaxie@gmail.com"
	this.Data["Author"] = "astaxie@gmail.com"

}

func (this *MainController) Explore() {

	this.Data["SiteName"] = "微信采集服务"
	this.Data["PageTitle"] = "浏览"
	this.Data["Keywords"] = "astaxie@gmail.com"
	this.Data["Description"] = "astaxie@gmail.com"
	this.Data["Author"] = "astaxie@gmail.com"
	this.TplName = "explore.html"
}
