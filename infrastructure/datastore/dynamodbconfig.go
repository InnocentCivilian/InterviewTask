package infrastructure

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/aws/aws-sdk-go/aws/session"
)

//returns correct dynamodb config based on running environment
func Database() *dynamodb.DynamoDB {
	switch os.Getenv("RUNNING_ENVIRONMENT") {
	case "Production":
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

		return dynamodb.New(sess)
	case "Local":
		sess := session.Must(session.NewSession())

		return dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://dynamodb:8000")})
	default:
		sess := session.Must(session.NewSession())

		return dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://dynamodb:8000")})
	}
}
