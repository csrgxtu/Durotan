package routers

import (
	"Durotan/controllers"
	"github.com/astaxie/beego"
)

func init() {
    // beego.Router("/", &controllers.MainController{})

    beego.Router("/auth/login", &controllers.AuthController{}, "get:Login")
}
