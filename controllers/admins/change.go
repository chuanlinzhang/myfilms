package admins



type ChangeController struct {
	AdminController
}

func (thhis *ChangeController) Get() {

}
func (this *ChangeController) Change()  {

}

func (this *ChangeController) Change1()  {
	admins:=this.GetSession("admins")
	this.Data["admins"]=admins
   this.TplName="admins-function1.html"
}
func (this *ChangeController) Change2()  {
	admins:=this.GetSession("admins")
	this.Data["admins"]=admins
	this.TplName="admins-function2.html"
}
func (this *ChangeController) Change3()  {
	admins:=this.GetSession("admins")
	this.Data["admins"]=admins
	this.TplName="admins-function3.html"
}

func (this *ChangeController) ChangeCustomers1()  {
	admins:=this.GetSession("admins")
	this.Data["admins"]=admins
	this.TplName="admins-member1.html"
}
func (this *ChangeController) ChangeCustomers2()  {
	admins:=this.GetSession("admins")
	this.Data["admins"]=admins
	this.TplName="admins-member2.html"
}