package users

/*
会员
 */
import (
	_ "gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"bs/myfilms/models/films"
	"time"
	"github.com/astaxie/beego/validation"
	"gopkg.in/mgo.v2/bson"
)

type Customers struct {
	LoginName string  `bson:"login_name"` //登陆账号
	Pwd       string  `bson:"pwd"`        //登录密码
	Email     string  `bson:"email"`      //Email
	Nickname  string  `bson:"nick_name"`  //昵称
	Name      string  `bson:"name"`       //会员实名
	Mobile    string  `bson:"mobile"`     //手机号码
	Sex       string  `bson:"sex"`        //性别
	Spent     float64 `bson:spent`        //总消费金额
	Balance   float64 `bson:"balance"`    //余额
	Memlv     int32   `bson:"menlv"`      //会员等级
	Available bool    `bson:"available"`  //资料是否有效
	Latest    string  `bson:"latest"`     //最后一次更新时间
}

var collectionCustomers *mgo.Collection

func init() {
	collectionCustomers = films.DB.C("customers")
}
func Register(loginName, pwd, email, nickName, name, moile, sex string) bool {
	customers := &Customers{
		LoginName: loginName,
		Pwd:       pwd,
		Email:     email,
		Nickname:  nickName,
		Name:      name,
		Mobile:    moile,
		Sex:       sex,
		Spent:     0,
		Balance:   0,
		Memlv:     1,
		Available: true,
		Latest:    time.Now().Format("2006-01-02 15:04:05"),
	}
	//验证用户输入信息格式是否正确
	valid:=validation.Validation{}
	_,err:=valid.Valid(customers)
	if err!=nil{
		return false
	}
	err = collectionCustomers.Insert(customers)
	if err != nil {
		return false
	}
	return true
}
func Login(loginName string) *Customers {
	customers:=&Customers{}
	err:=collectionCustomers.Find(bson.M{"login_name":loginName}).One(customers)
	if err!=nil{
		return nil
	}
	return customers
}