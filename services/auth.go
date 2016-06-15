package services

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"Durotan/models"
  "errors"
)

var UserCollection = beego.AppConfig.String("UserCollection")

func Login(mobile, password string) (err error, rtv models.User) {
  if CheckAndReconnect() != nil {
    return
  }

  var criteria = bson.M{"status": "visable", "mobile_number": mobile}
  err = Session.DB(DB).C(UserCollection).Find(criteria).One(&rtv)
  if err == nil {
    // beego.Info(rtv.Password)
    // beego.Info(GenerateGetMD5Password(mobile, password))
    if rtv.Password != GenerateGetMD5Password(mobile, password) {
      err = errors.New("账户信息错误")
      return
    }

    err, _ := CreateKongAPIConsumer(rtv.Password)
    if err != nil {
      beego.Info(err)
    }
  }

  return
}
