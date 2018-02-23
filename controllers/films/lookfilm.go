package films

type LookfilmController struct {
	FilmsController
}

func (this *LookfilmController)Get()  {

}
func (this *LookfilmController) Lookflim1()  {
	customers:=this.GetSession("customers")
	this.Data["customers"]=customers
 this.TplName="customers-ranking1.html"
}
func (this *LookfilmController) Lookflim2()  {
	customers:=this.GetSession("customers")
	this.Data["customers"]=customers
	this.TplName="customers-ranking2.html"
}
func (this *LookfilmController) Lookflim3()  {
	customers:=this.GetSession("customers")
	this.Data["customers"]=customers
	this.TplName="customers-ranking3.html"
}
