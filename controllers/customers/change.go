package customers

type ChangeController struct {
	CustomersController
}

func (this *ChangeController) Get() {

}
func (this *ChangeController) Change()  {

}

func (this *ChangeController) Change1()  {
	customers:=this.GetSession("customers")
   this.Data["customers"]=customers

	this.TplName="customers-function1.html"
}
func (this *ChangeController) Change2()  {
	customers:=this.GetSession("customers")
	this.Data["customers"]=customers
	this.TplName="customers-function2.html"
}
func (this *ChangeController) Change3()  {
	customers:=this.GetSession("customers")
	this.Data["customers"]=customers
	this.TplName="customers-function3.html"
}
func (this *ChangeController) Change4()  {
	customers:=this.GetSession("customers")
	this.Data["customers"]=customers
	this.TplName="customers-function4.html"
}