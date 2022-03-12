package infrastructure

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/aws/aws-sdk-go/aws/session"
)

func Database() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession())
	switch os.Getenv("RUNNING_ENVIRONMENT") {
	case "Local":
		return dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://dynamodb:8000")})
	default:
		return nil
	}
}
