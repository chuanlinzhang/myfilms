package tickets

type ChangeTicketController struct {
	TicketsController
}

func (this *ChangeTicketController) Get() {

}
func (this *ChangeTicketController) ChangeTicket() {

}
func (this *ChangeTicketController) ChangeTicket1() {
	this.TplName = "customers-tickets1.html"
}
func (this *ChangeTicketController) ChangeTicket2() {
	this.TplName = "customers-tickets2.html"
}
func (this *ChangeTicketController) ChangeTicket3() {
	this.TplName = "customers-tickets3.html"
}
