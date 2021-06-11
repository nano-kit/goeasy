# gate

The *gate* is an API Gateway.

## Supported Protocols

- HTTP
- WebSocket

## Features

- Translate HTTP round-trip to RPC to backend micro-services.
- Translate WebSocket to bidirectional streaming RPC to backend services.
- JSON payload is sent to and recv from backend services as it is.
- Backend services are discovered by service name and selected evenly across multiple instances.
- HTTP paths are dynamically mapped to services by the resolver as below.

## Resolver

*gate* dynamically routes to services using a namespace value and the HTTP path.
If a service have a name (io.goeasy.service.comet) and a method (Comet.Subscribe), let's first define some concepts:

```
io.goeasy.service.comet/Comet.Subscribe
<---1---> <--2--> <-3-> <------5------>
<----------4---------->
```

- Part 1 is a *namespace*
- Part 2 is a *type*
- Part 3 is an *alias*, or *short name*
- Altogether they compose part 4, which is *FQDN*, the global unique service name
- Part 5 is a *method* or *endpoint*

URLs are resolved as follows:

|       Path       |          Service          | Method  |
| ---------------- | ------------------------- | ------- |
| /foo/bar         | io.goeasy.service.foo     | Foo.Bar |
| /foo/bar/baz     | io.goeasy.service.foo     | Bar.Baz |
| /foo/bar/baz/cat | io.goeasy.service.foo.bar | Baz.Cat |

Versioned API URLs can easily be mapped to service names:

|      Path       |         Service          | Method  |
| --------------- | ------------------------ | ------- |
| /foo/bar        | io.goeasy.service.foo    | Foo.Bar |
| /v1/foo/bar     | io.goeasy.service.v1.foo | Foo.Bar |
| /v1/foo/bar/baz | io.goeasy.service.v1.foo | Bar.Baz |
| /v2/foo/bar     | io.goeasy.service.v2.foo | Foo.Bar |
| /v2/foo/bar/baz | io.goeasy.service.v2.foo | Bar.Baz |
