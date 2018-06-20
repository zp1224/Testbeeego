package routers
import (
	"pond/controllers"
	"github.com/astaxie/beego"
)

func init()  {
	beego.Router("/", &controllers.LoginController{}, "*:Index")
	beego.Router("/index2.html", &controllers.LoginController{}, "*:Index")
	beego.Router("/login", &controllers.LoginController{}, "*:Login")
	beego.Router("/logout", &controllers.LoginController{}, "*:Logout")
	//beego.Router("/getlogin", &controllers.LoginController{}, "*:GetLogin")
}

