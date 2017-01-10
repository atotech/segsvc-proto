all: segstore.pb.go segstore.pb.gw.go segstore.swagger.json

segstore.pb.go: segstore.proto
	protoc -I. -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --gogoslick_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. segstore.proto

segstore.pb.gw.go: segstore.proto
	protoc -I. -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. segstore.proto

segstore.swagger.json: segstore.proto
	protoc -I. -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. segstore.proto
