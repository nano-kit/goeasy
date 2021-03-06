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
- Built-in observability of metrics and tracing.
- Provides utilities such as image placeholder and file uploader out of box.
- HTTP paths are dynamically mapped to services by the resolver as below.

## API Resolver

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

## Web Resolver

In *goeasy* ecosystem, the *gate* is not only the API Gateway. As the *gate* is the only entrance for external requests, besides RESTful API, it also serves static web pages and HTTP GET (for example, an OAuth 2 Redirect URI or Callback URL).

Currently following paths are reserved by the resolver for static web pages:

* The root path `/`
* `/favicon.ico`
* `/portal`

HTTP GET can be handled by backend service. The query string is converted to JSON by [qson](https://github.com/nano-kit/go-micro/tree/main/util/qson).

## Monitoring

The *gate* exposes Prometheus text-based format metrics at `/metrics`.

## Image Placeholder

The *gate* generates custom placeholder images on the fly at `/placeholder`.

You can use the images in your HTML or CSS, like this:

```html
<img src="/placeholder?w=640&h=360">
```

## File Uploader

The *gate* can save artifacts.

You can upload files with `POST /o/upload`. The filename is taken from the original file. Or you can use `PUT /o/upload/(filename)`. In this case, the original file name is ignored, and the name is taken from the URL. Either case the URL to get the uploaded file is returned.

You can download files with `GET /o/(filename)`.

Add `/o/upload/*` to rules to allow only authenticated uploads.
