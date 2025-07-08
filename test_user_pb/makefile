
.PHONY: get
get:
	go get google.golang.org/grpc
	go get google.golang.org/protobuf

.PHONY: proto
proto:
	protoc -I proto proto/user.proto --go_out=./gen --go_opt=paths=source_relative \
	--go-grpc_out=./gen --go-grpc_opt=paths=source_relative

# ..._opt=paths=source_relative - опускает "option go_package" в proto файле

.PHONY: server
server:
	cmd /C start cmd /k "go run ./cmd/server/main.go"

.PHONY: client
client:
	cmd /C start cmd /k "go run ./cmd/client/main.go"