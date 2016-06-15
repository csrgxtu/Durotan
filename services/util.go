package services

import (
  "crypto/md5"
  "encoding/hex"
)

/**
 * 生成或者获取md5的密码，从mobile和明文的password
 */
func GenerateGetMD5Password(mobile, password string) string {
  hasher := md5.New()
  hasher.Write([]byte(mobile + password))
  return hex.EncodeToString(hasher.Sum(nil))
}
