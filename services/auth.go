package services

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"

	"Durotan/models"
)

var UserCollection = beego.AppConfig.String("UserCollection")

func Login(username, password string) (err error, rtv models.User) {
  if CheckAndReconnect() != nil {
    return
  }

  var criteria = bson.M{"status": "visable", "username": username}
  err = Session.DB(DB).C(UserCollection).Find(criteria).One(&rtv)
  if err == nil {
    beego.Info(rtv.UserName)
    beego.Info(rtv.Password)
  }
  return
}
