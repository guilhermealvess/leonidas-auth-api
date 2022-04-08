package service

import (
	"api-auth/src/adapter/grpc/pb"
	"context"
)

func (p *ApiServerServices) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Hello: "Hello" + in.Hello}, nil
}
