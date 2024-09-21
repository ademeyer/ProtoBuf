genproto:
	go mod init github.com/ademeyer/Protocol-Buffer
	protoc --go_out=./contact --go-grpc_out=./contact  --go_opt=paths=source_relative ProtoFile/*.proto
	go mod tidy
	go get github.com/brianvoe/gofakeit/v6   
	go run .