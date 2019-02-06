# grpcntt

gRPC network testing tools

## Using

Basically you start a server somewhere and ping it from somewhere else.

In the server box:

```sh
grpcntt serve --address :5555
```

In the client box:

```sh
grpcntt ping --address $SERVER_IP:5555
```

See `grpcntt help` output for reading about all available options.

## Building

If you don't need to change protobuf definitions then just running `go build` in
the root of the repository will produce the binary `grpcntt`.

Otherwise you need to have:

- [protoc](https://developers.google.com/protocol-buffers/docs/downloads.html)
- [protoc-gen-go](https://github.com/golang/protobuf/protoc-gen-go)
- [task](https://github.com/go-task/task)

Having that you can generate protobuf code as needed and build the binary by
running `task build`.

Which will produce the same binary.
