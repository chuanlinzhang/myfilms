package films

import (
	"gopkg.in/mgo.v2"
	"bs/myfilms/models/mongoDB"
)

/*
订单信息
 */
type Orders struct {
	OrderNo       string  `bson:"order_no"`       //订单编号
	OrderDate     string  `bson:"order_date"`     //下单时间
	OrderIdentify string  `bson:"order_identify"` //取票码
	Name          string  `bson:"name"`           //会员实名
	Amount        float64 `bson:"amount"`         //订单金额
	Mobile        string  `bson:"mobile"`         //付款方式
	Status        string  `bson:"status"`         //订单状态
	NameCh        string  `bson:"name_ch"`        //中文片名
	StartTime     string  `bson:"start_time"`     //放映时间
	Seat          string  `bson:"seat"`           //座位号
	Available     string  `bson:"availble"`       //资料是否有效
	Latest        string  `bson:"latest"`         //最后一次更新时间
}
var collectionOrders *mgo.Collection
func init() {
	dbname:=Dbname
	conn:=mongoDB.Dbsession.Copy()
	DB:=conn.DB(dbname)

	collectionOrders=DB.C("orders")
}
