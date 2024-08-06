proto:
	protoc --go_out=. --go-grpc_out=. ./api/proto/user.proto
tidy:
	go mod tidy
run:
	go run ./auth-server/cmd/server/main.go
test:
	go test ./auth-server/cmd/tests/*.go
.PHONY: proto tidy run test