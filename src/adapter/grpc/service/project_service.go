package service

import (
	"api-auth/src/adapter/grpc/pb"
	"api-auth/src/adapter/repository"
	"api-auth/src/entity"
	"api-auth/src/usecase"
	"context"
)

type ProjectServiceGRPC struct {
	Repository entity.ProjectRepository
	pb.UnimplementedProjectsServer
}

func NewProjectServiceGRPC(db repository.DocumentDB, cache repository.Cache) *ProjectServiceGRPC {
	repo := repository.NewProjectRepositoryDB(db, cache)

	return &ProjectServiceGRPC{
		Repository: repo,
	}
}

func (s *ProjectServiceGRPC) CreateProject(ctx context.Context, in *pb.CreateProjectRequest) (*pb.CreateProjectReply, error) {
	input := usecase.ProjectDtoInput{
		Name:          in.Project.Name,
		Description:   in.Project.Description,
		HashAlgorithm: in.Project.HashAlgorithm,
		RoundHash:     uint(in.Project.RoundHash),
	}

	processProject := usecase.NewProcessProject(s.Repository)

	output, err := processProject.ExecuteCreateNewProject(input)

	return &pb.CreateProjectReply{
		Error:      output.Error,
		StatusCode: output.Status,
		ProjectId:  output.ID,
		Credential: output.Credential,
		Key:        output.Key,
	}, err
}
