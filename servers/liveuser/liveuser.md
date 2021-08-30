# API Protocol

Table of Contents

* [Service User](#service-user)
    * [Method User.AddUser](#method-useradduser)
    * [Method User.QueryUser](#method-userqueryuser)
* [Service Wx](#service-wx)
    * [Method Wx.Login](#method-wxlogin)
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

客户端调用 wx.login() 获取临时登录凭证 code ，用此接口回传到开发者服务器。

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| code | string | 用户登录凭证（有效期五分钟）。开发者需要在开发者服务器后台调用 auth.code2Session， 使用 code 换取 openid、unionid、session_key 等信息 |

Response is empty





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

