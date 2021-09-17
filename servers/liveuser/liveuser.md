# API Protocol

Table of Contents

* [Service User](#service-user)
    * [Method User.Set](#method-userset)
    * [Method User.Get](#method-userget)
* [Service Wx](#service-wx)
    * [Method Wx.Login](#method-wxlogin)
    * [Method Wx.RenewToken](#method-wxrenewtoken)
    * [Method Wx.Prepay](#method-wxprepay)
    * [Method Wx.Postpay](#method-wxpostpay)
* [Service Order](#service-order)
    * [Method Order.Create](#method-ordercreate)
    * [Method Order.List](#method-orderlist)
* [Enums](#enums)
    * [Enum OrderRecord.State](#enum-orderrecordstate)
* [Objects](#objects)
    * [Object UserRecord](#object-userrecord)
    * [Object OrderProduct](#object-orderproduct)
    * [Object OrderRecord](#object-orderrecord)




## Service User



### Method User.Set

> POST /liveuser/User/Set <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

新增或更新自己的用户信息

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| user | [object UserRecord](#object-userrecord) | 需要更新的用户信息。其中 uid, update_at 可以不填。 |

Response is empty


### Method User.Get

> POST /liveuser/User/Get <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

获取自己的用户信息

Request is empty

Response parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| user | [object UserRecord](#object-userrecord) |  |





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


### Method Wx.Prepay

> POST /liveuser/Wx/Prepay <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

下单。调用该接口在微信支付服务后台生成预支付交易单，返回正确的预支付交易会话标识。

Request is empty

Response is empty


### Method Wx.Postpay

> POST /liveuser/Wx/Postpay <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

支付通知。微信支付通过支付通知接口将用户支付成功消息通知给开发者服务器。

Request is empty

Response is empty





## Service Order



### Method Order.Create

> POST /liveuser/Order/Create <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

创建订单

Request is empty

Response is empty


### Method Order.List

> POST /liveuser/Order/List <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

查询自己的订单

Request is empty

Response is empty





## Enums

### enum OrderRecord.State



Constants

|   Value   |   Name    |  Description |
| --------- | --------- | ------------ |
| 0  | CREATED |  |
| 1  | PAID |  |


## Objects

### object UserRecord

UserRecord 是用户信息

Attributes

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| uid | string | 唯一ID |
| name | string | 姓名 |
| agent | string | 终端 |
| update_at | int64 | 更新时间（毫秒时间戳） |
| avatar | string | 头像 |


### object OrderProduct



Attributes

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| order_id | uint64 |  |
| product_id | string |  |
| name | string |  |
| price | int32 |  |
| count | int32 |  |
| product_snapshot | uint64 |  |
| detail | string |  |


### object OrderRecord



Attributes

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| id | uint64 |  |
| uid | string |  |
| state | [enum OrderRecord.State](#enum-orderrecordstate) |  |
| amount | int32 |  |
| discount | int32 |  |
| pay | int32 |  |
| pay_at | int64 |  |
| products | array of [object OrderProduct](#object-orderproduct) |  |
| created_at | int64 |  |
| updated_at | int64 |  |
| deleted_at | int64 |  |

