# API Protocol

Table of Contents

* [Service User](#service-user)
    * [Method User.AddUser](#method-useradduser)
    * [Method User.QueryUser](#method-userqueryuser)
* [Enums](#enums)
* [Objects](#objects)
    * [Object UserRecord](#object-userrecord)




## Service User



### Method User.AddUser

> POST /liveuser/User/AddUser <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>



Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| user | [object UserRecord](#object-userrecord) |  |

Response is empty


### Method User.QueryUser

> POST /liveuser/User/QueryUser <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>



Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| uids | array of string |  |

Response parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| users | array of [object UserRecord](#object-userrecord) |  |





## Enums

## Objects

### object UserRecord



Attributes

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| uid | string | 唯一ID |
| name | string | 姓名 |
| agent | string | 终端 |
| email | string | 联系方式 |
| avatar | string | 头像 |

