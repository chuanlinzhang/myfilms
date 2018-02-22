package orders

type ChangeOrderController struct {
	OrdersController
}

func (this *ChangeOrderController) Get() {

}
func (this *ChangeOrderController) ChangeOrder()  {
	
}
func (this *ChangeOrderController) ChangeOrder1()  {
this.TplName="customers-order1.html"
}
func (this *ChangeOrderController) ChangeOrder2()  {
	this.TplName="customers-order2.html"
}