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

	return &pb.CreateProjectReply{
		Error:      output.Error,
		Success:    output.Success,
		ProjectId:  output.ID,
		Credential: output.Credential,
		Key:        output.Key,
	}, err
}
