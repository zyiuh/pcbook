gen:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/*.proto

clean:
	rm pb/*.go

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

test:
	go test -cover -race ./...