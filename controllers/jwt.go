package controllers

import (
	"github.com/astaxie/beego"

	"Durotan/models"
	"Durotan/services"
)

type JWTController struct {
	beego.Controller
}

func (this *JWTController) Login() {
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

func (this *JWTController) OrgLogin() {
  var rt models.Auth
  var email = this.GetString(":email")
  var password = this.GetString(":password")

  err, apikey, orgid := services.OrgLogin(email, password)
  if err != nil {
    rt.Msg = "创建apikey失败"
  } else {
    rt.Key = apikey
    rt.UserID = orgid
  }

  this.Data["json"] = &rt
  this.ServeJSON()
}
