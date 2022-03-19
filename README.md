# go-codebuild-aws-sdk


In this example project I am using codebuild aws sdk API where the goal is to get builds and their detailed deployment

The aws sdk version used in this example is old, there is a new version to use but it does not have the `*iface` interface and this lacks difficulty in unit testing implementations.

See more at:

- codebuildiface: https://pkg.go.dev/github.com/aws/aws-sdk-go-v2@v0.12.0/service/codebuild/codebuildiface
- other iface sample: https://aws.amazon.com/pt/blogs/developer/mocking-out-then-aws-sdk-for-go-for-unit-testing/
- Mocking out new API? : https://github.com/aws/aws-sdk-go-v2/issues/70 | https://github.com/aws/aws-sdk-go-v2/issues/50

# How to run

You need to have an aws webservice account and create a project on code build (see: https://aws.amazon.com/pt/codebuild/). To test you can create a plan without artifact.

After that, take your credential and enter this snippet code:

```go
sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("YOUR_ID", "YOUR_SECRET", "YOUR_TOKEN_OPTIONAL")},
	)
```

Done! run the command inside de path `go run main.go` and see your result. To run the tests : `go test`

