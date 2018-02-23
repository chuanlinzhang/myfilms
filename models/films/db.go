package films

import (

	"github.com/astaxie/beego"

)
var Dbname string
func init() {
	Dbname=beego.AppConfig.String("mongodb::dbname")


}
