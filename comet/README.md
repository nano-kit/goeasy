# comet

订阅和广播通道

Deprecated: consider using [NATS](https://github.com/nats-io/nats-server) 2.3 and [nats.ws](https://github.com/nats-io/nats.ws) for an end to end solution.

*comet* is a special service. It has a streaming endpoint *Subscribe*.
A client (usually a web browser such as Google Chrome) subscribes to the server by setting up a WebSocket connection to this streaming endpoint.

Once something happens in the server, a notification event is pushed to the client. The event should not be considered reliable. It is *at most once*.

The *comet* API is designed to be simple enough. The rules to keep the persistent connection active under unpredictable mobile networks:

- On receiving *heartbeat-probe*, send a *heartbeat-keepalive*
- On receiving nothing for a long time, re-connect
- On error, re-connect with reasonable backoff

## Overview

ENDPOINT (WebSocket)
    `/comet/subscribe`

REQUEST (*first-message*)
```
    {
        "type":"AUTH",
        "auth": { "token": "" },
        "join": { "rid": "" }
    }
```

REQUEST (*heartbeat-keepalive*)
```
    { "type":"HB" }
```

REQUEST (*join-room*)
```
    { "type":"JOIN", "join": { "rid": "" } }
```

RESPONSE (*heartbeat-probe*)
```
    { "type":"HB" }
```

RESPONSE (*server-pushed-event*)
```
    { "type":"PUSH", "push": { "evt": "" } }
```

## Details

See [comet.md](comet.md)
