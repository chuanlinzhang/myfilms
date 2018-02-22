package customers

type ChangeController struct {
	CustomersController
}

func (this *ChangeController) Get() {

}
func (this *ChangeController) Change()  {

}

func (this *ChangeController) Change1()  {
	this.TplName="customers-function1.html"
}
func (this *ChangeController) Change2()  {
	this.TplName="customers-function2.html"
}
func (this *ChangeController) Change3()  {
	this.TplName="customers-function3.html"
}
func (this *ChangeController) Change4()  {
	this.TplName="customers-function4.html"
}