package service

import (
	"api-auth/src/adapter/grpc/pb"
	"context"
)

type PingService struct {
	pb.UnimplementedHelloServer
}

func (p *PingService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Hello: "Hello" + in.Hello}, nil
}

func NewPingService() *PingService {
	return &PingService{}
}
