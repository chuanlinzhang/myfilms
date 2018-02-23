package films

import (
	"gopkg.in/mgo.v2"
	"bs/myfilms/models/mongoDB"
)

/*
影票信息
 */
type Tickets struct {
	Tno        string  `bson:"tno"`         //影票编号
	Zno        string  `bson:"zno"`         //放映厅编号
	NameCh     string  `bson:"name_ch"`     //中文片名
	NameEn     string  `bson:"name_en"`     //英文片名
	Amount     float64 `bson:"amount"`      //订单金额
	Mobile     string  `bson:"mobile"`      //付款方式
	Status     string  `bson:"status"`      //订单状态
	StartTime  string  `bson:"start_time"`  //放映时间
	FilmLength int32   `bson:"film_length"` //片长
	Seat       string  `bson:"seat"`        //座位号
	Available  string  `bson:"availble"`    //资料是否有效
	Latest     string  `bson:"latest"`      //最后一次更新时间
}
var collectionTickets *mgo.Collection
func init() {
	dbname:=Dbname
	conn:=mongoDB.Dbsession.Copy()
	DB:=conn.DB(dbname)

	collectionTickets=DB.C("tickets")
}
