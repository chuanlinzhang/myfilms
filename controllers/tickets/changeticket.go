package tickets

type ChangeTicketController struct {
	TicketsController
}

func (this *ChangeTicketController) Get() {

}
func (this *ChangeTicketController) ChangeTicket() {

}
func (this *ChangeTicketController) ChangeTicket1() {
	customers:=this.GetSession("customers")
	this.Data["customers"]=customers
	this.TplName = "customers-tickets1.html"
}
func (this *ChangeTicketController) ChangeTicket2() {
	customers:=this.GetSession("customers")
	this.Data["customers"]=customers
	this.TplName = "customers-tickets2.html"
}
func (this *ChangeTicketController) ChangeTicket3() {
	customers:=this.GetSession("customers")
	this.Data["customers"]=customers
	this.TplName = "customers-tickets3.html"
}
