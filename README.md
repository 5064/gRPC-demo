# gRPC-poc

Generate protobuffer from .proto file
```
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
$ go get -u google.golang.org/grpc
$ protoc -I=./api --go_out=plugins=grpc:./api ./api/language.proto
```

Run server (required running test database)
```
$ go run server/main.go
```

Send request
```
$ go run client/main.go 'japan'
$ go run client/main.go 'china'
$ go run client/main.go 'united states'
```