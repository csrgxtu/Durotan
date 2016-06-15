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

func (this *AuthController) Read() {
	skip, _ := this.GetInt(":skip")
	limit, _ := this.GetInt(":limit")
	var rt models.Result

	err, rtv := services.Read(skip, limit)
	if err != nil {
		rt.Error = true
		rt.Msg = "Failure"
		rt.Data = make([]models.Recs, 1)
		rt.Data[0] = err.Error()
	} else {
		rt.Error = false
		rt.Msg = "Successful"
		rt.Data = make([]models.Recs, len(rtv))
		for ix, value := range rtv {
			rt.Data[ix] = value
		}
	}

	this.Data["json"] = &rt
	this.ServeJSON()
}
