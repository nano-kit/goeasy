# comet

订阅和广播通道

*comet* is a special service. It has a streaming endpoint *Subscribe*.
A client (usually a web browser such as Google Chrome) subscribes to the server by setting up a WebSocket connection to this streaming endpoint.

Once something happens in the server, a notification event is pushed to the client. The event should not be considered reliable. It is *at most once*.

## API
