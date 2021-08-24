proto: protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
server: go run cmd/server/server.go
client: go run cmd/client/client.go
