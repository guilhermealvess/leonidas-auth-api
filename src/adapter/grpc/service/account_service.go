package service

import (
	"api-auth/src/adapter/grpc/pb"
	"api-auth/src/adapter/repository"
	"api-auth/src/entity"
	"api-auth/src/usecase"
	"context"
)

type AccountServiceGRPC struct {
	AccountRepository entity.AccountRepository
	ProjectRepository entity.ProjectRepository
	pb.UnimplementedAccountServicesServer
}

func NewAccountServiceGRPC(db repository.DocumentDB, cache repository.Cache) *AccountServiceGRPC {
	accountRepo := repository.NewAccountRepositoryDB(db, cache)
	projectRepo := repository.NewProjectRepositoryDB(db, cache)

	return &AccountServiceGRPC{
		AccountRepository: accountRepo,
		ProjectRepository: projectRepo,
	}
}

func (s *AccountServiceGRPC) CreateAccount(ctx context.Context, in *pb.CreateAccounttRequest) (*pb.CreateAccountReply, error) {
	input := usecase.AccountDtoInput{
		Name:       in.Account.Name,
		LastName:   in.Account.LastName,
		Credential: in.Credential,
		Key:        in.Key,
		Email:      in.Account.Email,
		Password:   in.Account.Password,
	}

	processAccount := usecase.NewProcessAccount(s.AccountRepository, s.ProjectRepository)
	output, err := processAccount.ExecuteCreateNewAccount(input)

	return &pb.CreateAccountReply{
		Error:     output.Error,
		Success:   output.Success,
		AccountId: output.ID,
	}, err
}
