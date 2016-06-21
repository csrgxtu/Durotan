package controllers

import (
	"github.com/astaxie/beego"

	"Durotan/models"
	"Durotan/services"
)

type KeyAuthController struct {
	beego.Controller
}

func (this *KeyAuthController) Login() {
  var rt models.Result
  var data models.Auth
  var mobile = this.GetString(":mobile")
  var password = this.GetString(":password")

  err, apikey, userid := services.Login(mobile, password)
  if err != nil {
    rt.Msg = "创建token失败"
    this.Ctx.ResponseWriter.WriteHeader(500)
  } else {
    rt.Msg = "创建token成功"
    data.Token = apikey
    data.AccountID = userid
  }

  this.Data["json"] = &rt
  this.ServeJSON()
}

func (this *KeyAuthController) OrgLogin() {
  var rt models.Result
  var data models.Auth
  var email = this.GetString(":email")
  var password = this.GetString(":password")

  err, apikey, orgid := services.OrgLogin(email, password)
  if err != nil {
    rt.Msg = "创建token失败"
    this.Ctx.ResponseWriter.WriteHeader(500)
  } else {
    rt.Msg = "创建token成功"
    data.Token = apikey
    data.AccountID = orgid
  }

  this.Data["json"] = &rt
  this.ServeJSON()
}
