package service

import (
	"api-auth/src/adapter/grpc/pb"
	"api-auth/src/usecase"
	"context"
)

func (s *ApiServerServices) CreateAccount(ctx context.Context, in *pb.CreateAccounttRequest) (*pb.CreateAccountReply, error) {
	input := usecase.AccountDtoInput{
		FirstName:             in.Account.FirstName,
		LastName:              in.Account.LastName,
		ApiKey:                in.ApiKey,
		Email:                 in.Account.Email,
		Username:              in.Account.Username,
		Password:              in.Account.Password,
		UrlRedirectActivation: in.UrlRedirectActivation,
	}

	processAccount := usecase.NewProcessAccount(s.AccountRepository, s.ProjectRepository)
	output, err := processAccount.ExecuteCreateNewAccount(input)

	if err != nil {
		return &pb.CreateAccountReply{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &pb.CreateAccountReply{
		Success:        true,
		Error:          "",
		AccountId:      output.ID,
		ActivationLink: output.Link,
	}, nil
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
