package controllers

type ErrorController struct {
	baseController
}

func (this *ErrorController) Error404() {
	this.Data["PageTitle"] = "404"
	this.Data["content"] = "page not found"
	this.Layout = "layout/_site_layout.html"
	this.TplName = "404.html"
}

func (this *ErrorController) Error501() {
	this.Data["PageTitle"] = "501"
	this.Data["content"] = "server error"
	this.Layout = "layout/_site_layout.html"
	this.TplName = "501.html"
}

func (this *ErrorController) Error500() {
	this.Data["PageTitle"] = "500"
	this.Data["content"] = "server error"
	this.Layout = "layout/_site_layout.html"
	this.TplName = "500.html"
}

func (this *ErrorController) ErrorDb() {
	this.Data["PageTitle"] = "DB"
	this.Data["content"] = "database is now down"
	this.Layout = "layout/_site_layout.html"
	this.TplName = "dberror.html"
}
