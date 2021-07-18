# API Protocol

Table of Contents

* [Enums](#enums)
* [Objects](#objects)
    * [Object Event](#object-event)
    * [Object Message](#object-message)




## Enums

## Objects

### object Event

Event is the asynchronous message using Go-Micro's message broker interface. refer to Comet.onEvent

Attributes

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| uid | string | The unique user identity |
| rid | string | The room identity |
| evt | string | The server-sent event goes through the downlink to the client. |


### object Message

Message is the general content predefined by comet for server-sent event. It is stringified to JSON and then assigned to `string evt`.

Attributes

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| server | string | A optional string identifying the server that generates event. |
| event | string | A string identifying the type of event. |
| data | string | The data field for the message. |
| id | string | The event ID. |
| time | int64 | The time this event happens. It is the number of nanoseconds elapsed since January 1, 1970 UTC. |

