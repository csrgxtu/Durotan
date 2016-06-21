package services

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"Durotan/models"
  "errors"
)

var JwtUserCollection = beego.AppConfig.String("UserCollection")

func JWTLogin(mobile, password string) (err error, token, aid string) {
  if CheckAndReconnect() != nil {
    return
  }

  var Account models.User
  var criteria = bson.M{"status": "visable", "mobile_number": mobile}
  err = Session.DB(DB).C(UserCollection).Find(criteria).One(&Account)
  if err == nil {
    if Account.Password != GenerateGetMD5Password(mobile, password) {
      err = errors.New("Unauthorized")
      return
    }

    aid = Account.ID.Hex()

    DeleteConsumer(Account.UserID)

    err, _ := CreateKongAPIConsumer(Account.UserID)
    if err != nil {
      err = errors.New("Server Internal Error")
    }

    err, token := GetKongJWTToken(Account.UserID)
    beego.Info(token)
  } else {
    err = errors.New("Server Internal Error")
  }

  return
}
