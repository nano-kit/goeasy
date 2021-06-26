# API Protocol

Table of Contents

* [Service Sequence](#service-sequence)
    * [Method Sequence.Next](#method-sequencenext)
    * [Method Sequence.Max](#method-sequencemax)
* [Enums](#enums)
* [Objects](#objects)




## Service Sequence

Sequence服务生成单调递增的序列号。每个命名都有自己的序列号。 命名一般可以是用户ID或者房间ID，意思是关于这个用户/房间的所 有的由序列号索引的消息。

### Method Sequence.Next

> POST /sequence/Sequence/Next <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

对请求的序列号，加1之后返回。每个命名的序列号，从1开始。

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| name | string | 序列号命名。每个命名有自己的递增序列号，与其它命名互不影响。 |

Response parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| value | uint64 | 序列号的值。 |


### Method Sequence.Max

> POST /sequence/Sequence/Max <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

对请求的序列号，返回已分配的最大值。序列号用于终端和后台 的数据同步。考虑到一种情况：终端由于特殊原因收不到新消息 通知，因此终端决定对后台做轮询。当终端来后台收取未读消息 时，Max可以作为第一道检查，避免直接访问消息缓存查找是否 有未读消息。

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| name | string | 序列号命名。 |

Response parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| value | uint64 | 序列号的值。 |





## Enums

## Objects
