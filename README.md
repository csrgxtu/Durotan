### Durotan
微服务，登陆认证模块，用于issue token。Durotan, Durotan, Durotan!!!

![Durotan](http://vignette3.wikia.nocookie.net/warcraftfilms/images/5/5f/Durotan1.jpg/revision/latest/scale-to-width-down/300?cb=20160718100822)

### 域名
|类型|域名|端口|
|:----:|:----:|:----:|
|测试|https://dev-durotan.beautifulreading.com|9002|
|生产|https://durotan.beautifulreading.com|9002|

### How to start
Durotan 是使用Golang, Beego实现的基本认证服务。

#### 启动
```bash
# clone the code
git clone https://github.com/beautifulreading/Durotan.git
cd Durotan


# start
bee run
```

#### 访问
```bash
http GET http://localhost:9002/
HTTP/1.1 200 OK
Access-Control-Allow-Methods: GET,HEAD,PUT,POST,DELETE
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 57
Content-Type: application/json; charset=utf-8
Date: Sun, 02 Oct 2016 05:41:47 GMT

{
    "code": 200,
    "data": "F* Off, I dont know ya!",
    "msg": "^_^"
}
```

#### 接口文档查看
[API Documentation](https://github.com/beautifulreading/Durotan/blob/master/doc/API.md)
