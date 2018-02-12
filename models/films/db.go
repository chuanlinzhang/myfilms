package films

import (
	"bs/myfilms/models/mongoDB"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)
var DB *mgo.Database
func init() {
	dbname:=beego.AppConfig.String("mongodb::dbname")
	conn:=mongoDB.Dbsession.Copy()
	DB=conn.DB(dbname)

}
