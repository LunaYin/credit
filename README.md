### install protocol buffer compiler and go plugins

Follow the instructions in the [quick start](https://grpc.io/docs/languages/go/quickstart/#prerequisites) document

### compile proto files

compile domain.proto file into go file
```
protoc --go_out=./persistence --go_opt=paths=source_relative --proto_path=./persistence persistence/domain.proto
```

compile credit.proto file into go files
```
protoc --go_out=./entity --go_opt=paths=source_relative --go-grpc_out=./entity --go-grpc_opt=paths=source_relative --proto_path=../protobuf/frontend  --proto_path=./entity entity/credit.proto
```

### start the server
```
source env.sh
go run credit.go
```

start the proxy-core project

### credit service call example

For user 'icn'
```
curl -s localhost:9000/credit/icn
curl -X POST -s -H "Content-Type: application/json" -d '{"quantity":10}' http://localhost:9000/credit/icn/add
```
