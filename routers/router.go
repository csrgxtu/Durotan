package routers

import (
	"Durotan/controllers"
	"github.com/astaxie/beego"
)

func init() {
  beego.Router("/", &controllers.DurotanController{}, "get:Welcome")
  beego.Router("/durotan/welcome", &controllers.DurotanController{}, "get:Welcome")

  // jwt
  beego.Router("/durotan/jwt", &controllers.JWTController{}, "get:Login")
  
  // // Key Auth
  // beego.Router("/auth/keyauth/login/:mobile/:password", &controllers.KeyAuthController{}, "get:Login")
  // beego.Router("/auth/keyauth/orglogin/:mobile/:password", &controllers.KeyAuthController{}, "get:OrgLogin")
  //
  // // JWT Auth
  // beego.Router("/auth/jwt/login/:mobile/:password", &controllers.JWTController{}, "get:Login")
  // beego.Router("/auth/jwt/orglogin/:email/:password", &controllers.JWTController{}, "get:OrgLogin")
}
