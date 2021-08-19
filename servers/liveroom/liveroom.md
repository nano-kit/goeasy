# API Protocol

Table of Contents

* [Service Room](#service-room)
    * [Method Room.Enter](#method-roomenter)
    * [Method Room.Send](#method-roomsend)
    * [Method Room.Recv](#method-roomrecv)
    * [Method Room.Leave](#method-roomleave)
* [Enums](#enums)
    * [Enum RoomMessage.Type](#enum-roommessagetype)
* [Objects](#objects)
    * [Object RoomMessage](#object-roommessage)
    * [Object MsgEnterRoom](#object-msgenterroom)
    * [Object MsgLeaveRoom](#object-msgleaveroom)
    * [Object MsgPlainText](#object-msgplaintext)




## Service Room

聊天室服务

### Method Room.Enter

> POST /liveroom/Room/Enter <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

进入聊天室

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| room | string | 聊天室ID |

Response parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| uids | array of string | 聊天室里在线的人列表 |


### Method Room.Send

> POST /liveroom/Room/Send <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

在聊天室里，发送消息

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| room | string | 聊天室ID |
| text | string | 消息内容 |

Response is empty


### Method Room.Recv

> POST /liveroom/Room/Recv <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

收取聊天室里的未读消息。 这是一个 long-polling 机制的方法，需要指定 `Request-Timeout` 头， 表示当聊天室里没有未读消息时，轮询多少秒。当聊天室里有未读消息时，立刻返回。

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| room | string | 聊天室ID |
| last_seq | uint64 | 客户端记住的已经收取到的最后一条消息的序列号 |

Response parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| msgs | array of [object RoomMessage](#object-roommessage) | 本次收取的所有未读消息，按seq排序，最小的seq必须比last_seq大 |


### Method Room.Leave

> POST /liveroom/Room/Leave <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

退出聊天室

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| room | string | 聊天室ID |

Response is empty





## Enums

### enum RoomMessage.Type

消息类型

Constants

|   Value   |   Name    |  Description |
| --------- | --------- | ------------ |
| 0  | UNSPECIFIED |  |
| 1  | ENTER_ROOM | 进入聊天室 |
| 2  | LEAVE_ROOM | 退出聊天室 |
| 3  | PLAIN_TEXT | 文本消息 |


## Objects

### object RoomMessage

聊天室消息

Attributes

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| room | string | // 聊天室ID |
| seq | uint64 | 聊天室里的每条消息都有唯一的seq，新消息的seq总是更大 |
| type | [enum RoomMessage.Type](#enum-roommessagetype) | 消息类型 |
| uid | string | 谁发出的消息 |
| send_at | int64 | 何时发出的消息（毫秒时间戳） |
| enter_room | [object MsgEnterRoom](#object-msgenterroom) | 具体消息对象，与消息类型对应 |
| leave_room | [object MsgLeaveRoom](#object-msgleaveroom) |  |
| plain_text | [object MsgPlainText](#object-msgplaintext) |  |


### object MsgEnterRoom

进入聊天室

It has no attributes


### object MsgLeaveRoom

退出聊天室

It has no attributes


### object MsgPlainText

聊天室文本消息

Attributes

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| text | string | 消息内容 |

