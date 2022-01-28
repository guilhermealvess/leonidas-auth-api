mkdir src/adapter/grpc/pb
protoc --proto_path=./src/adapter/grpc/proto --go_out=src/adapter/grpc/pb --go-grpc_out=src/adapter/grpc/pb src/adapter/grpc/proto/*.proto