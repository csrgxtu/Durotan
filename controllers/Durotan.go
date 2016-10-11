package controllers

import (
	"github.com/astaxie/beego"

	"Durotan/models"
)

type DurotanController struct {
	beego.Controller
}

func (this *DurotanController) Welcome() {
  var rt models.Result

  rt.Msg = "Durotan, 基本认证服务"

  this.Data["json"] = &rt
  this.ServeJSON()
}
