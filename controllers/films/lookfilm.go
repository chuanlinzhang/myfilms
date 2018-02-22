package films

type LookfilmController struct {
	FilmsController
}

func (this *LookfilmController)Get()  {

}
func (this *LookfilmController) Lookflim1()  {
 this.TplName="customers-ranking1.html"
}
func (this *LookfilmController) Lookflim2()  {
	this.TplName="customers-ranking2.html"
}
func (this *LookfilmController) Lookflim3()  {
	this.TplName="customers-ranking3.html"
}
