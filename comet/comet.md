# API Protocol

Table of Contents

* [Service Comet](#service-comet)
    * [Method Comet.Subscribe](#method-cometsubscribe)
    * [Method Comet.Publish](#method-cometpublish)
    * [Method Comet.Broadcast](#method-cometbroadcast)
    * [Method Comet.DumpSession](#method-cometdumpsession)
* [Enums](#enums)
    * [Enum MsgType](#enum-msgtype)
* [Objects](#objects)
    * [Object Room](#object-room)
    * [Object Session](#object-session)
    * [Object Heartbeat](#object-heartbeat)
    * [Object Auth](#object-auth)
    * [Object JoinRoom](#object-joinroom)
    * [Object ServerPush](#object-serverpush)
    * [Object Event](#object-event)




## Service Comet

Comet is a landing place where clients are continuously subscribing for downlink messages by a websocket or grpc connection.

### Method Comet.Subscribe

WebSocket bidirectional-streaming

> GET /comet/Comet/Subscribe <br/>

Subscribe for downlink messages while also sending uplink messages.

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| t | [enum MsgType](#enum-msgtype) | t is used to differentiate what this uplink message is |
| hb | [object Heartbeat](#object-heartbeat) |  |
| auth | [object Auth](#object-auth) |  |
| join | [object JoinRoom](#object-joinroom) |  |

Response parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| t | [enum MsgType](#enum-msgtype) | t is used to differentiate what this downlink message is |
| hb | [object Heartbeat](#object-heartbeat) |  |
| push | [object ServerPush](#object-serverpush) |  |


### Method Comet.Publish

> POST /comet/Comet/Publish <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

Publish an event to a specific client who is landing on this comet instance. Publish or Broadcast is not reliable innately! The event may not reach the client in case of comet restarts or broken connection. So it is not necessary to retry automatically. Usually you should design a SYNC protocol between client and server to achieve the reliable broadcasting.

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| uid | string | The unique user identity. Please note if an end-user has multiple terminals such as a phone, a pad, and a desktop, every terminal should use a unique identity if these devices are treated differently. |
| evt | string | The server-sent event goes through the downlink to the client. Its content is opaque, which means what is published here reaches client unmodified. Only the concrete business can explain this event. |

Response is empty


### Method Comet.Broadcast

> POST /comet/Comet/Broadcast <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

Broadcast an event to all the clients in a specific room on this comet instance.

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| rid | string | The room identity, default (empty) to do a world broadcast, which means all the clients on this comet instance are published with the event. Otherwise, the broadcast is delivered to the clients within the room. |
| evt | string | The server-sent event goes through the downlink to the client. Its content is opaque, which means what is published here reaches client unmodified. Only the concrete business can explain this event. |

Response is empty


### Method Comet.DumpSession

> POST /comet/Comet/DumpSession <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

Dump all the clients' session on this comet instance. It is a debugging method.

Request is empty

Response parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| world | array of [object Session](#object-session) | All the sessions in the world in this comet instance |
| rooms | array of [object Room](#object-room) | All the rooms in this comet instance |





## Enums

### enum MsgType



Constants

|   Value   |   Name    |  Description |
| --------- | --------- | ------------ |
| 0  | HB | Heartbeat is sent on downlink and uplink, to keep the persistent connection alive |
| 1  | AUTH | Auth is sent on uplink as the first message for Comet.Subscribe |
| 2  | JOIN | JoinRoom is sent on uplink to join the specified room. JOIN can be sent together with AUTH, which means connect to comet and join the specified room immediately. A client follows only one room at a time. If a client is already in a room, JOIN a different room implies quitting the last room where the client stays; JOIN the same room is a no-op. A connected client is always considered in the world in spite of maybe in a room. |
| 3  | PUSH | ServerPush is sent on downlink to push event to client |


## Objects

### object Room

Room is a (virtual) place where session gathers.

Attributes

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| rid | string | The room identity |
| room | array of [object Session](#object-session) | All the sessions in this room. |


### object Session

Session is a user's session. It is created when a client subscribes to this comet, and destroyed when the client disconnected.

Attributes

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| uid | string | The unique user identity |
| rid | string | The room identity |
| birth | string | When is the session created |


### object Heartbeat



Attributes

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |


### object Auth



Attributes

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| token | string | Usually a JWT is used. Comet.Subscribe extracts user's account information from this token. |


### object JoinRoom



Attributes

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| rid | string | The room identity, default (empty) to quit any room and stay only in the world. |


### object ServerPush



Attributes

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| evt | string | The server-sent event goes through the downlink to the client. Its content is opaque, which means what is published reaches client unmodified. Only the concrete business can explain this event. |


### object Event

Event is the asynchronous message using Go-Micro's message broker interface. refer to Comet.onEvent

Attributes

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| uid | string | The unique user identity |
| rid | string | The room identity |
| evt | string | The server-sent event goes through the downlink to the client. |

