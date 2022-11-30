proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=paths=source_relative:. pkg/pb/*.proto

run-server:
	@go run cmd/main.go