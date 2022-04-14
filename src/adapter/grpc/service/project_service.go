package service

import (
	"api-auth/src/adapter/grpc/pb"
	"api-auth/src/usecase"
	"context"
)

func (s *ApiServerServices) CreateProject(ctx context.Context, in *pb.CreateProjectRequest) (*pb.CreateProjectReply, error) {
	input := usecase.ProjectDtoInput{
		Name:          in.Project.Name,
		Description:   in.Project.Description,
		HashAlgorithm: in.Project.HashAlgorithm,
		RoundHash:     uint(in.Project.RoundHash),
	}

	processProject := usecase.NewProcessProject(s.ProjectRepository)
	output, err := processProject.ExecuteCreateNewProject(input)

	if err != nil {
		return &pb.CreateProjectReply{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &pb.CreateProjectReply{
		Success: true,
		ApiKey:  output.ApiKey,
	}, nil
}
