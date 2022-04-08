package service

import (
	"api-auth/src/adapter/grpc/pb"
	"api-auth/src/usecase"
	"context"
)

func (s *ApiServerServices) CreateAccount(ctx context.Context, in *pb.CreateAccounttRequest) (*pb.CreateAccountReply, error) {
	input := usecase.AccountDtoInput{
		Name:                  in.Account.Name,
		LastName:              in.Account.LastName,
		ApiKey:                in.ApiKey,
		Email:                 in.Account.Email,
		Password:              in.Account.Password,
		UrlRedirectActivation: in.UrlRedirectActivation,
	}

	processAccount := usecase.NewProcessAccount(s.AccountRepository, s.ProjectRepository)
	output, err := processAccount.ExecuteCreateNewAccount(input)

	return &pb.CreateAccountReply{
		Error:          output.Error,
		Success:        output.Success,
		AccountId:      output.ID,
		ActivationLink: output.Link,
	}, err
}

func (s *ApiServerServices) RefreshActivationLinkAccount(ctx context.Context, in *pb.RefreshActivationLinkAccountRequest) (*pb.RefreshActivationLinkAccountResponse, error) {
	input := in.Id

	processAccount := usecase.NewProcessAccount(s.AccountRepository, s.ProjectRepository)

	link, err := processAccount.RefreshActivationLinkAccount(input)
	sucess := true
	e := ""

	if err != nil {
		sucess = false
		e = err.Error()
	}

	return &pb.RefreshActivationLinkAccountResponse{
		Success: sucess,
		Error:   e,
		Link:    link,
	}, err
}
