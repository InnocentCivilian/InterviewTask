package service

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// DynamoDBRepository -
type DynamoDBRepository struct {
	session   *dynamodb.DynamoDB
	tableName string
}

// NewDynamoDBRepository -
func NewDynamoDBRepository(ddb *dynamodb.DynamoDB, tableName string) *DynamoDBRepository {
	return &DynamoDBRepository{ddb, tableName}
}

// Get a user
func (r *DynamoDBRepository) Get(ctx context.Context, id string) (*Device, error) {
	device := &Device{}
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	result, err := r.session.GetItemWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	if err := dynamodbattribute.UnmarshalMap(result.Item, &device); err != nil {
		return nil, err
	}

	return device, nil
}

// Create a user
func (r *DynamoDBRepository) Create(ctx context.Context, device *Device) error {
	item, err := dynamodbattribute.MarshalMap(device)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(r.tableName),
	}
	_, err = r.session.PutItemWithContext(ctx, input)
	return err
}
