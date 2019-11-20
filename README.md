# gRPC-poc

generate protobuffer from .proto file
```
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
$ go get -u google.golang.org/grpc
$ protoc -I=./api --go_out=plugins=grpc:./api ./api/language.proto
```

