package controllers

import (
	"github.com/astaxie/beego"

	"Durotan/models"
	"Durotan/services"
)

type AuthController struct {
	beego.Controller
}

func (this *AuthController) Login() {
  var rt models.Auth
  var mobile = this.GetString(":mobile")
  var password = this.GetString(":password")

  err, apikey, userid := services.Login(mobile, password)
  if err != nil {
    rt.Msg = "创建apikey失败"
  } else {
    rt.Key = apikey
    rt.UserID = userid
  }

  this.Data["json"] = &rt
  this.ServeJSON()
}
