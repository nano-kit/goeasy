# goeasy

Build the web real-time messaging systems in lightning speed.

快速打造您的web实时通讯体系。

Goeasy is a cloud service and solution for Go developers that makes it incredibly simple to add real-time web functionality to your applications. What is "real-time web" functionality? It's the ability to have your server-side code push content to the connected clients as it happens, in real-time.

## What can it be used for?

Pushing data from the server to the client (not just browser clients) has always been a tough problem. Goeasy makes it dead easy and handles all the heavy lifting for you.

## Design considerations

Goeasy is complete and production ready. It differs with other similar solutions on these unique perspective.

### Interface

Goeasy is a cloud service. It provides the real-time ability by exposing HTTP APIs. The API specification is described with Protocol Buffers (proto3). There is a convention which translates HTTP API endpoints to protobuf Service RPCs. So, the API is documented, consistent and extensible. Nearly every client programming language works without too much effort.

The ability of pushing data from the server to the client requires every connected client maintains a persistent connection with server. WebSocket is used in this case. At the edge of cloud, the WebSocket connection is translated to gRPC bidirectional streaming RPC.

HTTP long-polling is supported in client platform where WebSocket is thought to be complicated or inapplicable. Unlike WebSocket, which requires client is always connected to the server, so that the data can be delivered to the client, HTTP long-polling provides a reliable real-time notification mechanism that no message is missing even the client may lose its connection.

### Elasticity

Goeasy is broken into several servers. Each server do one thing well. One server could have multiple instances. Servers communicate with each other in protobuf Service RPCs.

Among these servers, there is a special kind of server that also exposes HTTP APIs. they are called *gate*. Client talks to goeasy through *gate*. The *gate*s act as the access layer and stay at the edge of goeasy cloud.

There are other kind of servers: *auth* does authentication and authorization; *comet* manages persistent connections; *room* provides a chat room service having reliable messaging; *user*  manages users' profiles. These servers construct a micro-service system that has great elasticity when deployed in cloud-native environment.

### Infrastructure

Goeasy does not reinvent the well-known infrastructure utilities. It utilizes cloud-native distributed solutions such as *CockroachDB* the distributed SQL database; *Cassandra* the distributed NoSQL database; *NATS* the cloud and edge native messaging system; *MinIO* the S3 compatible object storage. However SQLite and Redis are also good enough options for start-up solutions.

### Deployment

Goeasy could be deployed as bundle of servers, it could also be deployed as an all in one executable. This is convenient for testing and minimum viable product verification.

## Documents

* [Setup a SaaS](HowToSetup.md)
* [Project Layout](ProjectLayout.md)
* [Architecture](doc/index.md)
* [Reference: Go-Micro in Action](https://nano-kit.github.io/go-micro-in-action/)
