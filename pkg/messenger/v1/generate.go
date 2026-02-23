package v1

//go:generate protoc --proto_path=../../.. --go_out=../../.. --go_opt=module=messenger --go-grpc_out=../../.. --go-grpc_opt=module=messenger api/proto/messenger/v1/messenger.proto
