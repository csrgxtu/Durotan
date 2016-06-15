package services

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"Durotan/models"
)

var UserCollection = beego.AppConfig.String("UserCollection")

func Login(mobile, password string) (err error, rtv models.User) {
  if CheckAndReconnect() != nil {
    return
  }

  var criteria = bson.M{"status": "visable", "mobile_number": mobile}
  err = Session.DB(DB).C(UserCollection).Find(criteria).One(&rtv)
  if err == nil {
    beego.Info(rtv.Password)
    beego.Info(GenerateGetMD5Password(mobile, password))
  }
  return
}
