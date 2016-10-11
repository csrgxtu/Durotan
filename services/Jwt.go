package services

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"Durotan/models"
  "errors"
)

var JwtUserCollection = beego.AppConfig.String("UserCollection")

func JWTLogin(Type, openid, mobile, password string) (err error, token, aid string) {
  if CheckAndReconnect() != nil {
    return
  }

  var Account models.User
  if Type == "Wechat" {
    var criteria = bson.M{"status": "visable", "openid": openid}
    err = Session.DB(DB).C(UserCollection).Find(criteria).One(&Account)
    if err == nil {
      aid = Account.ID.Hex()
      token = GenerateJwtToken()
    }
  } else {
    var criteria = bson.M{"status": "visable", "mobile_number": mobile}
    err = Session.DB(DB).C(UserCollection).Find(criteria).One(&Account)
    if err == nil {
      if Account.Password != GenerateGetMD5Password(mobile, password) {
        err = errors.New("Account Unvalid")
        return
      }

      aid = Account.ID.Hex()
      token = GenerateJwtToken()
    }
  }

  return
}
