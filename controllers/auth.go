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
  var rt models.Result
  var username = this.GetString(":username")
  var password = this.GetString(":password")

  err, rtv := services.Login(username, password)
  if err != nil {
    rt.Error = true
    rt.Msg = err.Error()
  } else {
    rt.Error = false
    rt.Msg = "Successful"
    rt.Data = make([]models.Recs, 1)
    rt.Data[0] = &rtv
  }

  this.Data["json"] = &rt
  this.ServeJSON()
}
