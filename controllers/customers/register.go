package customers

import (
	"github.com/astaxie/beego"
	"bs/myfilms/models/users"

)

type RegisterController struct {
	CustomersController
}

func (this *RegisterController) Get() {
	this.TplName = "customers-register.html"
}
func (this *RegisterController) Register() {
	cname := this.GetString("cname")
	cpwd1 := this.GetString("cpwd1")
	cpwd2 := this.GetString("cpwd2")
	cemail := this.GetString("cemail")
	nick_name := this.GetString("nick_name")
	name := this.GetString("name")
	mobile := this.GetString("mobile")
	sex := this.GetString("sex")

	if cpwd1 != cpwd2 {
		this.Redirect(beego.URLFor("customers.RegisterController.Get"), 302)
		return
	}
   b:=users.Register(cname,cpwd1,cemail,nick_name,name,mobile,sex)
   if b==false{
   	this.Redirect(beego.URLFor("customers.RegisterController.Get"),302)
   	return
   }
   this.TplName="login-index.html"
}
