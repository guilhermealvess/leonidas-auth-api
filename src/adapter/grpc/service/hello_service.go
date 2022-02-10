package service

import (
	"api-auth/src/adapter/grpc/pb"
	"context"
)

type Ping struct {
	pb.UnimplementedHelloServer
}

func (p *Ping) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Hello: "hello"}, nil
}
