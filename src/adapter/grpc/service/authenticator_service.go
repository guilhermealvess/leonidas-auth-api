package service

import (
	"api-auth/src/adapter/grpc/pb"
	"api-auth/src/usecase"
	"context"
)

func (s *ApiServerServices) SignIn(ctx context.Context, in *pb.SigninRequest) (*pb.SigninReply, error) {
	input := usecase.ProcessSignInput{
		Credential: in.Credential,
		Key:        in.Key,
		Email:      in.Email,
		Password:   in.Password,
	}

	processProcessAuthenticator := usecase.NewProcessAuthenticator(s.ProjectRepository, s.AccountRepository, s.jwtMaker)
	output, err := processProcessAuthenticator.Sign(input)
	if err != nil {
		return &pb.SigninReply{
			Success: false,
			Error:   err.Error(),
			Token:   "",
		}, err
	}

	return &pb.SigninReply{
		Success: true,
		Token:   output.Token,
		Error:   "",
	}, nil

}

func (s *ApiServerServices) VerifyToken(ctx context.Context, in *pb.VerifyTokenRequest) (*pb.VerifyTokenReply, error) {
	input := usecase.ProcessVerifyTokenInput{
		Token:     in.Token,
		ProjectId: in.ProjectId,
	}

	processProcessAuthenticator := usecase.NewProcessAuthenticator(s.ProjectRepository, s.AccountRepository, s.jwtMaker)
	output, err := processProcessAuthenticator.VerifyToken(input)
	if err != nil {
		return &pb.VerifyTokenReply{
			Success: false,
			Error:   err.Error(),
			Payload: &pb.Payload{},
		}, nil
	}

	return &pb.VerifyTokenReply{
		Success: true,
		Error:   "",
		Payload: &pb.Payload{
			Id:        output.Payload.ID,
			Email:     output.Payload.Email,
			IssueAt:   output.Payload.IssuedAt.Unix(),
			ExpiredAt: output.Payload.ExpiredAt.Unix(),
		},
	}, nil
}
