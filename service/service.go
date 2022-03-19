package service

import (
	"github.com/aws/aws-sdk-go/service/codebuild"
	"github.com/aws/aws-sdk-go/service/codebuild/codebuildiface"
)

type CodeBuildService struct {
	Client codebuildiface.CodeBuildAPI
}

func (s CodeBuildService) BatchGetBuilds(params *codebuild.BatchGetBuildsInput) (*codebuild.BatchGetBuildsOutput, error) {
	result, err := s.Client.BatchGetBuilds(params)
	return result, err
}

func (s CodeBuildService) ListBuilds(param *codebuild.ListBuildsInput) (*codebuild.ListBuildsOutput, error) {
	result, err := s.Client.ListBuilds(param)
	return result, err
}
