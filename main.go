package main

import (
  "github.com/astaxie/beego/plugins/cors"
  "github.com/astaxie/beego"
	_ "Durotan/routers"
)

func main() {
  // CORS for https://foo.* origins, allowing:
  // - PUT and PATCH methods
  // - Origin header
 // // - Credentials share
  beego.InsertFilter("*", beego.BeforeRouter,cors.Allow(&cors.Options{
  AllowOrigins: []string{"*"},
  AllowMethods: []string{"*"},
  AllowHeaders: []string{"Origin"},
  ExposeHeaders: []string{"Content-Length"},
  AllowCredentials: true,
  }))

  
	beego.Run()
}
