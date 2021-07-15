# API Protocol

Table of Contents

* [Service LiveRoom](#service-liveroom)
    * [Method LiveRoom.Send](#method-liveroomsend)
    * [Method LiveRoom.Recv](#method-liveroomrecv)
* [Enums](#enums)
    * [Enum RoomMessage.Type](#enum-roommessagetype)
* [Objects](#objects)
    * [Object RoomMessage](#object-roommessage)
    * [Object MsgEnterRoom](#object-msgenterroom)
    * [Object MsgLeaveRoom](#object-msgleaveroom)
    * [Object MsgPlainText](#object-msgplaintext)




## Service LiveRoom

聊天室服务

### Method LiveRoom.Send

> POST /liveroom/LiveRoom/Send <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

在聊天室里，发送消息

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| room | string | 聊天室ID |
| msg | [object RoomMessage](#object-roommessage) | 本次发送的消息 |

Response is empty


### Method LiveRoom.Recv

> POST /liveroom/LiveRoom/Recv <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

收取聊天室里的未读消息

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| room | string | 聊天室ID |
| last_seq | uint64 | 客户端记住的已经收取到的最后一条消息的序列号 |

Response parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| msgs | array of [object RoomMessage](#object-roommessage) | 本次收取的所有未读消息，按seq排序，最小的seq必须比last_seq大 |





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
| seq | uint64 | 聊天室里的每条消息都有唯一的seq，新消息的seq总是更大 |
| typ | [enum RoomMessage.Type](#enum-roommessagetype) | 消息类型 |
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
