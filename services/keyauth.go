package services

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"Durotan/models"
  "errors"
)

var UserCollection = beego.AppConfig.String("UserCollection")

func KeyAuthLogin(mobile, password string) (err error, apikey, userid string) {
  if CheckAndReconnect() != nil {
    return
  }

  var user models.User
  var criteria = bson.M{"status": "visable", "mobile_number": mobile}
  err = Session.DB(DB).C(UserCollection).Find(criteria).One(&user)
  if err == nil {
    if user.Password != GenerateGetMD5Password(mobile, password) {
      err = errors.New("Unauthorized")
      return
    }

    userid = user.ID.Hex()

    DeleteConsumer(user.Password)

    err, _ := CreateKongAPIConsumer(user.Password)
    if err != nil {
      beego.Info(err)
    }

    err, apikey = GetKongKeyAuthToken(user.Password)
  } else {
    err = errors.New("Server Internal Error")
  }

  return
}

func OrgLogin(mobile, password string) (err error, apikey, orgid string) {
  if CheckAndReconnect() != nil {
    return
  }

  var org models.User
<<<<<<< HEAD
  var criteria = bson.M{"status": "visable", "mobile_number": mobile}
=======
  var criteria = bson.M{"status": "visable", "mobile_number": mobile, "organization": bson.M{"$exists": true}}
>>>>>>> 5fde1a149a67f3c085d9b12e3992620c2776e2c4
  err = Session.DB(DB).C(UserCollection).Find(criteria).One(&org)
  if err == nil {
    if org.Password != GenerateGetMD5Password(mobile, password) {
      err = errors.New("Unauthorized")
      return
    }

    orgid = org.ID.Hex()

    DeleteConsumer(org.Password)

    err, _ := CreateKongAPIConsumer(org.Password)
    if err != nil {
      beego.Info(err)
    }

    err, apikey = GetKongKeyAuthToken(org.Password)
  } else {
    err = errors.New("Server Internal Error")
  }

  return
}
