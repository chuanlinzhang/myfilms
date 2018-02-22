package routers

import (
	"bs/myfilms/controllers"
	"github.com/astaxie/beego"
	"bs/myfilms/controllers/customers"
	"bs/myfilms/controllers/films"
	"bs/myfilms/controllers/tickets"
	"bs/myfilms/controllers/orders"
	"bs/myfilms/controllers/admins"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.BaseController{})

	beego.Router("/customers/login", &customers.LoginController{})
	beego.Router("/customers/login/check", &customers.LoginController{}, "Post:Login")

	beego.Router("/customers/register", &customers.RegisterController{})
	beego.Router("/customers/register/check", &customers.RegisterController{}, "Post:Register")

	beego.Router("/customers-ranking1", &films.LookfilmController{}, "Get:Lookflim1")
	beego.Router("/customers-ranking2", &films.LookfilmController{}, "Get:Lookflim2")
	beego.Router("/customers-ranking3", &films.LookfilmController{}, "Get:Lookflim3")

	beego.Router("/customers-function1", &customers.ChangeController{}, "Get:Change1")
	beego.Router("/customers-function2", &customers.ChangeController{}, "Get:Change2")
	beego.Router("/customers-function3", &customers.ChangeController{}, "Get:Change3")
	beego.Router("/customers-function4", &customers.ChangeController{}, "Get:Change4")

	beego.Router("/customers-tickets1", &tickets.ChangeTicketController{}, "Get:ChangeTicket1")
	beego.Router("/customers-tickets2", &tickets.ChangeTicketController{}, "Get:ChangeTicket2")
	beego.Router("/customers-tickets3", &tickets.ChangeTicketController{}, "Get:ChangeTicket3")

	beego.Router("/customers-order1", &orders.ChangeOrderController{}, "Get:ChangeOrder1")
	beego.Router("/customers-order2", &orders.ChangeOrderController{}, "Get:ChangeOrder2")





	beego.Router("/admins/login", &admins.LoginController{})
	beego.Router("/admins/login/check", &admins.LoginController{}, "Post:Login")

	beego.Router("/admins/register", &admins.RegisterController{})
	beego.Router("/admins/register/check", &admins.RegisterController{}, "Post:Register")
}
