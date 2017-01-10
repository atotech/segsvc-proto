all: segsvc.pb.go segsvc.pb.gw.go segsvc.swagger.json

segsvc.pb.go: segsvc.proto
	protoc -I. -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --gogoslick_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. segsvc.proto

segsvc.pb.gw.go: segsvc.proto
	protoc -I. -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. segsvc.proto

segsvc.swagger.json: segsvc.proto
	protoc -I. -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. segsvc.proto
