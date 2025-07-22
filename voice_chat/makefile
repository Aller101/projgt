.PHONY: go_get
go_get:
	go get github.com/gorilla/websocket

.PHONY: run_server run_client
run_server:
	cmd /C start cmd /k "go run ./cmd/server/main.go"

run_client:
	cmd /C start cmd /k "go run ./cmd/client/main.go"