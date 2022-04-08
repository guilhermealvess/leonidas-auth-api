package service

import (
	"api-auth/src/adapter/grpc/pb"
	"api-auth/src/adapter/jwt"
	"api-auth/src/adapter/repository"
	"api-auth/src/entity"
)

type ApiServerServices struct {
	AccountRepository entity.AccountRepository
	ProjectRepository entity.ProjectRepository
	jwtMaker          jwt.JWT

	pb.UnimplementedApiV1ServicesServer
}

func NewAccountServiceGRPC(db repository.DocumentDB, cache repository.Cache, jwtMaker jwt.JWT) *ApiServerServices {
	accountRepo := repository.NewAccountRepositoryDB(db, cache)
	projectRepo := repository.NewProjectRepositoryDB(db, cache)

	return &ApiServerServices{
		AccountRepository: accountRepo,
		ProjectRepository: projectRepo,
		jwtMaker:          jwtMaker,
	}
}
