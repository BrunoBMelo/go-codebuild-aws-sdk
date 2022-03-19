package service_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/codebuild"
	"github.com/aws/aws-sdk-go/service/codebuild/codebuildiface"
	"github.com/stretchr/testify/assert"
)

var TestsTable = []struct {
	MockCodeBuildClient      *mockCodeBuildClient
	MockBatchGetBuildsOutput *codebuild.BatchGetBuildsOutput
	MockListBuildsOutput     *codebuild.ListBuildsOutput
}{
	{
		MockCodeBuildClient: &mockCodeBuildClient{},
		MockBatchGetBuildsOutput: &codebuild.BatchGetBuildsOutput{
			Builds: []*codebuild.Build{
				{Id: aws.String("mockCodebuild:7f3d1356-aa8b-46c9-bab0-6af725ed90b6"),
					BuildNumber:  aws.Int64(1),
					BuildStatus:  aws.String("SUCCESS"),
					ProjectName:  aws.String("mockCodeBuild"),
					CurrentPhase: aws.String("COMPLETED")},
			},
			BuildsNotFound: []*string{},
		},
		MockListBuildsOutput: &codebuild.ListBuildsOutput{
			Ids: []*string{
				aws.String("mockCodebuild:7f3d1356-aa8b-46c9-bab0-6af725ed90b6"),
			},
		},
	},
}

func TestCodeBuildService(t *testing.T) {

	for _, suiteTest := range TestsTable {

		names, _ := suiteTest.MockCodeBuildClient.ListBuilds(&codebuild.ListBuildsInput{
			SortOrder: aws.String("ASCENDING"),
		})

		assert.Equal(t, names, suiteTest.MockListBuildsOutput)

		builds, _ := suiteTest.MockCodeBuildClient.BatchGetBuilds(&codebuild.BatchGetBuildsInput{Ids: names.Ids})

		assert.Equal(t, builds, suiteTest.MockBatchGetBuildsOutput)
	}
}

type mockCodeBuildClient struct {
	codebuildiface.CodeBuildAPI
}

func (s mockCodeBuildClient) BatchGetBuilds(params *codebuild.BatchGetBuildsInput) (*codebuild.BatchGetBuildsOutput, error) {
	return &codebuild.BatchGetBuildsOutput{
		Builds: []*codebuild.Build{
			{Id: aws.String("mockCodebuild:7f3d1356-aa8b-46c9-bab0-6af725ed90b6"),
				BuildNumber:  aws.Int64(1),
				BuildStatus:  aws.String("SUCCESS"),
				ProjectName:  aws.String("mockCodeBuild"),
				CurrentPhase: aws.String("COMPLETED")},
		},
		BuildsNotFound: []*string{},
	}, nil
}

func (s mockCodeBuildClient) ListBuilds(param *codebuild.ListBuildsInput) (*codebuild.ListBuildsOutput, error) {
	return &codebuild.ListBuildsOutput{
		Ids: []*string{
			aws.String("mockCodebuild:7f3d1356-aa8b-46c9-bab0-6af725ed90b6"),
		},
	}, nil
}
