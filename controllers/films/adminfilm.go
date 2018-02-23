package films

type AdminfilmController struct {
	FilmsController
}

func (this *AdminfilmController)Get()  {

}
func (this *AdminfilmController) Adminflim1()  {
	admins:=this.GetSession("admins")
	this.Data["admins"]=admins
	this.TplName="admins-movie1.html"
}
func (this *AdminfilmController) Adminflim2()  {
	admins:=this.GetSession("admins")
	this.Data["admins"]=admins
	this.TplName="admins-movie2.html"
}
func (this *AdminfilmController) Adminflim3()  {
	admins:=this.GetSession("admins")
	this.Data["admins"]=admins
	this.TplName="admins-movie3.html"
}
func (this *AdminfilmController) Adminflim4()  {
	admins:=this.GetSession("admins")
	this.Data["admins"]=admins
	this.TplName="admins-movie4.html"
}