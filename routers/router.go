package routers

import (
	"Durotan/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/auth/login/:mobile/:password", &controllers.AuthController{}, "get:Login")
    beego.Router("/auth/orglogin/:email/:password", &controllers.AuthController{}, "get:OrgLogin")
}
