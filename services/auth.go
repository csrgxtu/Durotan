package services

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"Durotan/models"
  "errors"
)

var UserCollection = beego.AppConfig.String("UserCollection")

func Login(mobile, password string) (err error, apikey, userid string) {
  if CheckAndReconnect() != nil {
    return
  }

  var user models.User
  var criteria = bson.M{"status": "visable", "mobile_number": mobile}
  err = Session.DB(DB).C(UserCollection).Find(criteria).One(&user)
  if err == nil {
    // beego.Info(rtv.Password)
    // beego.Info(GenerateGetMD5Password(mobile, password))
    if user.Password != GenerateGetMD5Password(mobile, password) {
      err = errors.New("账户信息错误")
      return
    }

    userid = user.ID.Hex()

    DeleteConsumer(user.Password)

    err, _ := CreateKongAPIConsumer(user.Password)
    if err != nil {
      beego.Info(err)
    }

    err, apikey = GetKongAPIKey(user.Password)
  }

  return
}

func OrgLogin(email, password string) (err error, apikey, orgid string) {
  if CheckAndReconnect() != nil {
    return
  }

  var org models.User
  var criteria = bson.M{"status": "visable", "email": email}
  err = Session.DB(DB).C(UserCollection).Find(criteria).One(&org)
  if err == nil {
    // beego.Info(rtv.Password)
    // beego.Info(GenerateGetMD5Password(mobile, password))
    if org.Password != GenerateGetMD5Password(email, password) {
      err = errors.New("账户信息错误")
      return
    }

    orgid = org.ID.Hex()

    DeleteConsumer(org.Password)

    err, _ := CreateKongAPIConsumer(org.Password)
    if err != nil {
      beego.Info(err)
    }

    err, apikey = GetKongAPIKey(org.Password)
  }

  return
}
