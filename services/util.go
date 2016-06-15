package services

import (
  "github.com/parnurzeal/gorequest"
  "github.com/astaxie/beego"
  "crypto/md5"
  "encoding/hex"
  "errors"
)

/**
 * 生成或者获取md5的密码，从mobile和明文的password
 */
func GenerateGetMD5Password(mobile, password string) string {
  hasher := md5.New()
  hasher.Write([]byte(mobile + password))
  return hex.EncodeToString(hasher.Sum(nil))
}

func CreateKongAPIConsumer(consumer_id string) (err error, consumer string) {
  var ConsumerAPI = beego.AppConfig.String("KongConsumerAPI")

  request := gorequest.New()
  res, _, errs := request.Post(ConsumerAPI).
    Send(`username=` + consumer_id).
    End()
  if len(errs) != 0 {
    err = errs[0]
    return
  }

  if res.StatusCode != 201 {
    err = errors.New("创建key过程失败")
    return
  }

  consumer = consumer_id

  return
}

func GetKongAPIKey(consumer_id string) (err error, apikey string) {
  var ConsumerAPI = beego.AppConfig.String("KongConsumerAPI") + consumer_id + "/key-auth"

  request := gorequest.New()
  res, body, errs := request.Post(ConsumerAPI).
    End()
  if len(errs) != 0 {
    err = errs[0]
    return
  }

  if res.StatusCode != 201 {
    err = errors.New("创建key过程失败")
    return
  }

  beego.Info(body)

  return
}
