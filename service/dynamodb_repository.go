package service

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/innocentcivilian/interviewtask/util"
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
			"Id": {
				S: aws.String(id),
			},
		},
	}

	result, err := r.session.GetItemWithContext(ctx, input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, errors.New(util.NotFound)
	}
	if err := dynamodbattribute.UnmarshalMap(result.Item, &device); err != nil {
		return nil, err
	}

	return device, nil
}

// Create a device
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
