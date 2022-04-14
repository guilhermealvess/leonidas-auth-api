package service

import (
	"api-auth/src/adapter/grpc/pb"
	"api-auth/src/usecase"
	"context"
)

func (s *ApiServerServices) CreateAccount(ctx context.Context, in *pb.CreateAccounttRequest) (*pb.CreateAccountReply, error) {
	input := usecase.AccountDtoInput{
		FirstName: in.Account.FirstName,
		LastName:  in.Account.LastName,
		ApiKey:    in.ApiKey,
		Email:     in.Account.Email,
		Username:  in.Account.Username,
		Password:  in.Account.Password,
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
		Success:   true,
		AccountId: output.ID,
	}, nil
}

func (s *ApiServerServices) ActivateAccount(ctx context.Context, in *pb.ActivateAccountRequest) (*pb.ActivateAccountResponse, error) {
	input := usecase.ActivationAccountDtoInput{
		ApiKey:   in.ApiKey,
		Username: in.Username,
	}

	processAccount := usecase.NewProcessAccount(s.AccountRepository, s.ProjectRepository)
	if err := processAccount.ActivateAccount(input); err != nil {
		return &pb.ActivateAccountResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &pb.ActivateAccountResponse{
		Success: true,
	}, nil
}

func (s *ApiServerServices) SetNewPassowrd(ctx context.Context, in *pb.NewPassowrdRequest) (*pb.NewPassowrdResponse, error) {
	input := usecase.SetNewPassowrdDtoInput{
		ApiKey:   in.ApiKey,
		Username: in.Username,
	}
	processAccount := usecase.NewProcessAccount(s.AccountRepository, s.ProjectRepository)

	if err := processAccount.SetNewPassword(input); err != nil {
		return &pb.NewPassowrdResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &pb.NewPassowrdResponse{
		Success: true,
	}, nil
}
