genproto:
	protoc --go_out=./contact --go-grpc_out=./contact  --go_opt=paths=source_relative ProtoFile/*.proto
	go mod init github.com/ademeyer/Protocol-Buffer
	go mod tidy
	go get github.com/brianvoe/gofakeit/v6   
	go run .