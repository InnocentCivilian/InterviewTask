package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//this function will create database schema for device entity;dynamodb is schemaless and only needs keys to be defined.
func Migrate() {
	params := &dynamodb.CreateTableInput{
		TableName: aws.String("Devices"),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err := Database().CreateTable(params)
	if err != nil {
		panic("Device database migration failed")
	}
}
