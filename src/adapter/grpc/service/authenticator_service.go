package service

import (
	"api-auth/src/adapter/grpc/pb"
	"api-auth/src/entity"
	"api-auth/src/usecase"
	"context"
)

type AuthenticatorServiceGRPC struct {
	pb.UnimplementedAuthenticatorServer
	projectRepository entity.ProjectRepository
	accountRepository entity.AccountRepository
	jwtMaker          usecase.JWT
}

func NewAuthenticatorServiceGRPC(projectRepo entity.ProjectRepository, accountRepo entity.AccountRepository, jwtMaker usecase.JWT) *AuthenticatorServiceGRPC {
	return &AuthenticatorServiceGRPC{
		projectRepository: projectRepo,
		accountRepository: accountRepo,
		jwtMaker:          jwtMaker,
	}
}

func (s *AuthenticatorServiceGRPC) SignIn(ctx context.Context, in *pb.SigninRequest) (*pb.SigninReply, error) {
	input := usecase.ProcessSignInput{
		Credential: in.Credential,
		Key:        in.Key,
		Email:      in.Email,
		Password:   in.Password,
	}

	processProcessAuthenticator := usecase.NewProcessAuthenticator(s.projectRepository, s.accountRepository, s.jwtMaker)
	output, err := processProcessAuthenticator.Sign(input)
	if err != nil {
		return &pb.SigninReply{
			Error:      "",
			StatusCode: 500,
			Token:      "",
		}, err
	}

	return &pb.SigninReply{
		Error:      "200",
		StatusCode: 200,
		Token:      output.Token,
	}, nil

}

func (s *AuthenticatorServiceGRPC) VerifyToken(ctx context.Context, in *pb.VerifyTokenRequest) (*pb.VerifyTokenReply, error) {
	input := usecase.ProcessVerifyTokenInput{
		Token: in.Token,
	}

	processProcessAuthenticator := usecase.NewProcessAuthenticator(s.projectRepository, s.accountRepository, s.jwtMaker)
	output, err := processProcessAuthenticator.VerifyToken(input)
	if err != nil {
		return &pb.VerifyTokenReply{
			Error:      err.Error(),
			StatusCode: 500,
			Payload:    &pb.Payload{},
		}, err
	}

	return &pb.VerifyTokenReply{
		Error:      "",
		StatusCode: 200,
		Payload: &pb.Payload{
			Id:        output.Payload.ID,
			Email:     output.Payload.Email,
			IssueAt:   output.Payload.IssuedAt.Unix(),
			ExpiredAt: output.Payload.ExpiredAt.Unix(),
		},
	}, nil
}