package infrastructure

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/aws/aws-sdk-go/aws/session"
)

//returns correct dynamodb config based on running environment
func Database() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession())
	switch os.Getenv("RUNNING_ENVIRONMENT") {
	case "Production":
		return nil // todo : production enviroment db connection config to be added
	case "Local":
		return dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://dynamodb:8000")})
	default:
		return dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://dynamodb:8000")})
	}
}
