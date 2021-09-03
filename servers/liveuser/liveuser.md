# API Protocol

Table of Contents

* [Service User](#service-user)
    * [Method User.AddUser](#method-useradduser)
    * [Method User.QueryUser](#method-userqueryuser)
* [Service Wx](#service-wx)
    * [Method Wx.Login](#method-wxlogin)
    * [Method Wx.RenewToken](#method-wxrenewtoken)
* [Enums](#enums)
* [Objects](#objects)
    * [Object UserRecord](#object-userrecord)




## Service User



### Method User.AddUser

> POST /liveuser/User/AddUser <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

新增或更新用户信息

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| user | [object UserRecord](#object-userrecord) |  |

Response is empty


### Method User.QueryUser

> POST /liveuser/User/QueryUser <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

查询用户信息

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| uids | array of string |  |

Response parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| users | array of [object UserRecord](#object-userrecord) |  |





## Service Wx



### Method Wx.Login

> POST /liveuser/Wx/Login <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

客户端调用 wx.login() 获取临时登录凭证 code ，用此接口回传到开发者服务器。 开发者服务器处理之后，返回开发者服务器的自定义登录态。 关于自定义登录态的解释，可以参考 https://auth0.com/blog/refresh-tokens-what-are-they-and-when-to-use-them/

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| code | string | 用户登录凭证（有效期五分钟）。开发者需要在开发者服务器后台调用 auth.code2Session， 使用 code 换取 openid、unionid、session_key 等信息 |

Response parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| access_token | string | 该用户调用开发者服务器后台的凭据，用来识别用户身份 |
| refresh_token | string | 用来换取新的 access_token，客户端应该保存在本地存储 |
| expiry | int64 | access_token 凭证到期的时间，格式为Unix时间戳 |


### Method Wx.RenewToken

> POST /liveuser/Wx/RenewToken <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

开发者服务器的自定义登录态里的 access_token 到期之前，用此接口获取新的 access_token。

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| refresh_token | string | 客户端保存在本地存储的上次的 refresh_token |

Response parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| access_token | string | 该用户调用开发者服务器后台的凭据，用来识别用户身份 |
| refresh_token | string | 用来换取新的 access_token，客户端应该保存在本地存储。 取决于是否开启了 Refresh Token Rotation，它可能与请求时的 refresh_token 不同 |
| expiry | int64 | access_token 凭证到期的时间，格式为Unix时间戳 |





## Enums

## Objects

### object UserRecord

UserRecord 是用户信息

Attributes

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| uid | string | 唯一ID |
| name | string | 姓名 |
| agent | string | 终端 |
| update_at | string | 更新时间(秒) |
| avatar | string | 头像 |

