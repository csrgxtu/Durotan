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

  this.Data["json"] = &rt
  this.ServeJSON()
}
