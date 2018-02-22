package users

/*
管理员
 */
import (
	_ "gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"bs/myfilms/models/films"
	"github.com/astaxie/beego/validation"
)

type Admins struct {
	AdminLogin    string `bson:"admin_login"`    //管理员登陆账号
	AdminPwd      string `bson:"admin_pwd"`      //管理员登陆密码
	AdminNickName string `bson:"admin_nickname"` //管理员昵称
	AdminName     string `bson:"admin_name"`     //管理员实名
	AdminMobile   string `bson:"admin_mobile"`   //管理员手机号码
	AdminSex      string `bson:"admin_sex"`      //管理员性别
	AdminEmail    string `bson:"admin_email"`    //管理员Email
	Available     bool   `bson:"available"`      //资料是否有效
	Latest        string `bson:"latest"`         //最后一次更新时间
}

var collectionAdmins *mgo.Collection
func init() {
	collectionAdmins=films.DB.C("admins")
}

func RegisterA(adminLogin,adminPwd,adminNickname,adminName,adminMobile,adminSex,adminEmail string) bool {
	admins:=&Admins{
		AdminLogin:adminLogin,
		AdminPwd:adminPwd,
		AdminNickName:adminNickname,
		AdminName:adminName,
		AdminMobile:adminMobile,
		AdminSex:adminSex,
		AdminEmail:adminEmail,
	}
	//验证用户输入信息格式是否正确
	valid:=validation.Validation{}
	_,err:=valid.Valid(admins)
	if err!=nil{
		return false
	}
	err=collectionCustomers.Insert(admins)
	if err!=nil{
		return false
	}
	return true
}