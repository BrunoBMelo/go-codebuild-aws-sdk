package main

import (
	"fmt"

	"os"

	"bbmello/codebuild/service"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codebuild"
)

func main() {

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("YOUR_ID", "YOUR_SECRET", "YOUR_TOKEN_OPTIONAL")},
	)

	// Initialize a struct
	svc := service.CodeBuildService{
		Client: codebuild.New(sess),
	}

	// Get the list of builds
	names, err := svc.ListBuilds(&codebuild.ListBuildsInput{
		SortOrder: aws.String("ASCENDING"),
	})

	if err != nil {
		fmt.Println("Got error listing builds: ", err)
		os.Exit(1)
	}

	// Get information about each build
	builds, err := svc.BatchGetBuilds(&codebuild.BatchGetBuildsInput{Ids: names.Ids})

	if err != nil {
		fmt.Println("Got error getting builds: ", err)
		os.Exit(1)
	}

	for _, build := range builds.Builds {
		fmt.Printf("Build id: %s\n", aws.StringValue(build.Id))
		fmt.Printf("Project Name: %s\n", aws.StringValue(build.ProjectName))
		fmt.Printf("Current Phase:   %s\n", aws.StringValue(build.CurrentPhase))
		fmt.Printf("Build Status:  %s\n", aws.StringValue(build.BuildStatus))
		fmt.Printf("Build Number:  %d\n", *build.BuildNumber)
	}
}
