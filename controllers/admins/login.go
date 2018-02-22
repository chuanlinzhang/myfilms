package admins

import "fmt"

type LoginController struct {
	AdminController
}

func (this *LoginController)Get()  {
	this.TplName="admin-login.html"
}
func (this *LoginController)Login()  {
	admin_login:=this.GetString("admin_login")
	admin_pwd:=this.GetString("admin_pwd")

	fmt.Println(admin_login,admin_pwd,"++++++++++++++++++++++++++++++++++++++++")
}