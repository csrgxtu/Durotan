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
  var rt models.Result
  var data models.Auth
  var Type = this.GetString("type")
  var openid = this.GetString("openid")
  var mobile = this.GetString("mobile")
  var password = this.GetString("password")

  err, token, userid := services.JWTLogin(Type, openid, mobile, password)
  if err != nil {
    rt.Msg = "创建token失败"
    this.Ctx.ResponseWriter.WriteHeader(500)
  } else {
    rt.Msg = "创建token成功"
    data.Token = token
    data.AccountID = userid
    rt.Data = data
  }

  this.Data["json"] = &rt
  this.ServeJSON()
}

//
// func (this *JWTController) OrgLogin() {
//   var rt models.Result
//   var data models.Auth
//   var email = this.GetString(":email")
//   var password = this.GetString(":password")
//
//   err, apikey, orgid := services.OrgLogin(email, password)
//   if err != nil {
//     rt.Msg = "创建apikey失败"
//     this.Ctx.ResponseWriter.WriteHeader(500)
//   } else {
//     rt.Msg = "创建token成功"
//     data.Token = apikey
//     data.AccountID = orgid
//   }
//
//   this.Data["json"] = &rt
//   this.ServeJSON()
// }
