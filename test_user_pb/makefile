
.PHONY: get
get:
	go get google.golang.org/grpc
	go get google.golang.org/protobuf

.PHONY: proto
proto:
	protoc -I proto proto/user.proto --go_out=./gen --go_opt=paths=source_relative \
	--go-grpc_out=./gen --go-grpc_opt=paths=source_relative

# ..._opt=paths=source_relative - опускает "option go_package" в proto файле

UNAME_S := $(shell uname -s)

.PHONY: insecure_server insecure_client

insecure_server:
ifeq ($(OS),Windows_NT)
	@echo "Running server on Windows"
	cmd /C start cmd /k "go run ./cmd/insecure/server/main.go"
else ifeq ($(UNAME_S),Linux)
	@echo "Running server on Linux"
	gnome-terminal -- go run ./cmd/insecure/server/main.go
else ifeq ($(UNAME_S),Darwin)
	@echo "Running server on macOS"
	open -a Terminal go run ./cmd/insecure/server/main.go
else
	@echo "Unsupported OS: $(UNAME_S)"
endif

insecure_client:
ifeq ($(OS),Windows_NT)
	@echo "Running client on Windows"
	cmd /C start cmd /k "go run ./cmd/insecure/client/main.go"
else ifeq ($(UNAME_S),Linux)
	@echo "Running client on Linux"
	gnome-terminal -- go run ./cmd/insecure/client/main.go
else ifeq ($(UNAME_S),Darwin)
	@echo "Running client on macOS"
	open -a Terminal go run ./cmd/insecure/client/main.go
else
	@echo "Unsupported OS: $(UNAME_S)"
endif

.PHONY: secure_server
secure_server:
	cmd /C start cmd /k "go run ./cmd/secure/server/main.go"

.PHONY: secure_client
secure_client:
	cmd /C start cmd /k "go run ./cmd/secure/client/main.go"

.PHONY: openssl
openssl:
	openssl genrsa -out ca.key 2048
	openssl req -new -x509 -days 365 -key ca.key -subj "/C=CN/ST=GD/L=SZ/O=Acme, Inc./CN=Acme Root CA" -out ca.crt
	openssl req -newkey rsa:2048 -nodes -keyout server.key -subj "/C=CN/ST=GD/L=SZ/O=Acme, Inc./CN=localhost" -out server.csr
	echo subjectAltName=DNS:localhost > extfile.cnf
	openssl x509 -req -extfile extfile.cnf -days 365 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt
	del extfile.cnf
	