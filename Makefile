genproto:
	protoc --go_out=./contact --go-grpc_out=./contact  --go_opt=paths=source_relative ProtoFile/*.proto
