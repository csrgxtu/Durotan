package services

import (
  "github.com/parnurzeal/gorequest"
  jwt "github.com/dgrijalva/jwt-go"
  "github.com/astaxie/beego"
  "Durotan/models"
  "crypto/md5"
  "encoding/hex"
  "encoding/json"
  "errors"
  "time"
)


/**
 * GenerateJwtToken
 * 生成jwt token
 */
func GenerateJwtToken(aid string) string {
  var alg = beego.AppConfig.String("JwtAlg")
  var secret = beego.AppConfig.String("JwtSecret")
  var expire, _ = beego.AppConfig.Int64("JwtExpire")

  SigningKey := []byte(secret)

  // Create the Claims
  claims := &jwt.MapClaims{
    "aid": aid,
    "exp": time.Now().Add(time.Hour * time.Duration(expire)).Unix(), // token到期时间
  }
  // claims := &jwt.StandardClaims{
  //   "aid": aid,
  //   ExpiresAt: expire,
  // }

  token := jwt.NewWithClaims(jwt.GetSigningMethod(alg), claims)
  tokenStr, _ := token.SignedString(SigningKey)

  return tokenStr
}

/**
 * 生成或者获取md5的密码，从mobile和明文的password
 */
func GenerateGetMD5Password(mobile, password string) string {
  hasher := md5.New()
  hasher.Write([]byte(mobile + password))
  return hex.EncodeToString(hasher.Sum(nil))
}

func DeleteConsumer(consumer_id string) (err error) {
  var ConsumerAPI = beego.AppConfig.String("KongConsumerAPI") + consumer_id

  request := gorequest.New()
  _, _, errs := request.Delete(ConsumerAPI).
    End()
  if len(errs) != 0 {
    err = errs[0]
    return
  }

  // beego.Info(res.StatusCode)
  // beego.Info(body)

  return
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

func GetKongKeyAuthToken(consumer_id string) (err error, apikey string) {
  var ConsumerAPI = beego.AppConfig.String("KongConsumerAPI") + consumer_id + "/key-auth"
  // beego.Info(ConsumerAPI)

  request := gorequest.New()
  res, body, errs := request.Post(ConsumerAPI).
    // Set("Content-Type", "application/x-www-form-urlencoded").
    // Send("query=bicycle").
    Type("form").
    End()
  if len(errs) != 0 {
    err = errs[0]
    return
  }

  beego.Info(body)
  if res.StatusCode != 201 {
    err = errors.New("创建token过程失败")
    return
  }

  var keyAuthJson models.KeyAuthKey
  err = json.Unmarshal([]byte(body), &keyAuthJson)
  if err != nil {
    return
  }

  apikey = keyAuthJson.Key

  return
}

func GetKongJWTToken(consumer_id string) (err error, token string) {
  var ConsumerAPI = beego.AppConfig.String("KongConsumerAPI") + consumer_id + "/jwt"

  request := gorequest.New()
  res, body, errs := request.Post(ConsumerAPI).
    Type("form").
    End()
  if len(errs) != 0 {
    err = errs[0]
    return
  }

  beego.Info(body)
  if res.StatusCode != 201 {
    err = errors.New("创建token过程失败")
    return
  }

  return
}
