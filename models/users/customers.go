package users

/*
会员
 */
import (
	_ "gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"bs/myfilms/models/films"
)

type Customers struct {
	LoginName string  `bson:"login_name"` //登陆账号
	Pwd       string  `bson:"pwd"`        //登录密码
	Email     string  `bson:"email"`      //Email
	Nickname  string  `bson:"nick_name"`  //昵称
	Name      string  `bson:"name"`       //会员实名
	Mobile    string  `bson:"mobile"`     //手机号码
	Phone     string  `bson:"phone"`      //电话
	Sex       string  `bson:"sex"`        //性别
	Spent     float64 `bson:spent`        //总消费金额
	Balance   float64 `bson:"balance"`    //余额
	Memlv     string  `bson:"menlv"`      //会员等级
	Available bool    `bson:"available"`  //资料是否有效
	Latest    string  `bson:"latest"`     //最后一次更新时间
}
var collectionCustomers *mgo.Collection
func init() {
	collectionCustomers=films.DB.C("customers")
}
