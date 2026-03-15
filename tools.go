//go:build generate

package messenger

//go:generate protoc --proto_path=. --go_out=. --go_opt=module=github.com/aniats/messenger --go-grpc_out=. --go-grpc_opt=module=github.com/aniats/messenger api/proto/messenger/v1/messenger.proto
