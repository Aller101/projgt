
.PHONY: get
get:
	go get google.golang.org/grpc
	go get google.golang.org/protobuf

.PHONY: proto
proto:
	protoc -I proto proto/user.proto --go_out=./gen --go_opt=paths=source_relative \
	--go-grpc_out=./gen --go-grpc_opt=paths=source_relative

# ..._opt=paths=source_relative - опускает "option go_package" в proto файле

.PHONY: insecure_server
insecure_server:
	cmd /C start cmd /k "go run ./cmd/insecure/server/main.go"

.PHONY: insecure_client
insecure_client:
	cmd /C start cmd /k "go run ./cmd/insecure/client/main.go"