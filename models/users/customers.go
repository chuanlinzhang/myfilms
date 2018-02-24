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
	"bs/myfilms/models/mongoDB"
	"github.com/astaxie/beego/logs"


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
	dbname := films.Dbname
	conn := mongoDB.Dbsession.Copy()
	DB := conn.DB(dbname)
	collectionCustomers = DB.C("customers")
}
func Register(loginName, pwd, email, nickName, name, moile, sex string) bool {

	customers := &Customers{}
	err := collectionCustomers.Find(bson.M{"login_name": loginName}).One(customers)
	if err == nil {
		logs.Info("用户已存在")
		return false
	}
	customers = &Customers{
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
	valid := validation.Validation{}
	_, err = valid.Valid(customers)
	if err != nil {
		return false
	}
	err = collectionCustomers.Insert(customers)
	if err != nil {
		return false
	}
	return true
}
func Login(loginName string) *Customers {
	customers := &Customers{}
	err := collectionCustomers.Find(bson.M{"login_name": loginName}).One(customers)
	if err != nil {
		return nil
	}
	return customers
}
func Change(loginName, email, nickName, name, moile, sex string) bool {
	err := collectionCustomers.Update(bson.M{"login_name": loginName}, bson.M{"$set": bson.M{
		"email":     email,
		"nick_name": nickName,
		"name":      name,
		"mobile":    moile,
		"sex":       sex,
	}})
	if err != nil {
		return false
	}
	return true
}
func ChangePwd(loginName, newPwd1 string) bool {
	err := collectionCustomers.Update(bson.M{"login_name": loginName}, bson.M{"$set": bson.M{"pwd": newPwd1}})
	if err != nil {
		return false
	}
	return true
}
func TopUp(loginName,pwd string,balance float64) bool {
     err:=collectionCustomers.Update(bson.M{"login_name":loginName,"pwd":pwd},bson.M{"$set":bson.M{"balance":balance}})
     if err!=nil{
     	return false
	 }
	 return true
}
//获取所以的用户
func CustomersTreeGrid() []Customers  {
	//查询所有
	 list :=make([]Customers,0)//定义一个切片
	collectionCustomers.Find(nil).All(&list)
   return list
}

func DelCus(loginName string) bool {
	err:=collectionCustomers.Remove(bson.M{"login_name":loginName})
	if err!=nil{
		return false
	}
	return true
}
////将资源列表转成treegrid格式
//func resourceList2TreeGrid(list []*Customers) []*Customers {
//	result := make([]*Customers, 0)
//	for _, item := range list {
//		if item.Parent == nil || item.Parent.Id == 0 {
//			item.Level = 0
//			result = append(result, item)
//			result = resourceAddSons(item, list, result)
//		}
//	}
//	return result
//}