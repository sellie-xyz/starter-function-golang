package {{cookiecutter.module}}

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	log "github.com/sirupsen/logrus"
)

type DynamoDBRepository struct {
	session   *dynamodb.DynamoDB
	tableName string
}

func NewDynamoDBRepository(ddb *dynamodb.DynamoDB, tableName string) *DynamoDBRepository {
	return &DynamoDBRepository{ddb, tableName}
}

func (r *DynamoDBRepository) SaveObject(ctx context.Context, object *Object)  error {
	item, err := dynamodbattribute.MarshalMap(object)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(r.tableName),
	}
	_, err = r.session.PutItemWithContext(ctx, input)
	log.Info(err)
	return err
}

func (r *DynamoDBRepository) SaveStateTransition(ctx context.Context, statusTransitions *StatusTransitions) error {
	item, err := dynamodbattribute.MarshalMap(statusTransitions)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(r.tableName),
	}
	_, err = r.session.PutItemWithContext(ctx, input)
	log.Info(err)
	return err}

func (r *DynamoDBRepository) SaveCurrentState(ctx context.Context, currentStatus *CurrentStatus) error {
	item, err := dynamodbattribute.MarshalMap(currentStatus)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(r.tableName),
	}
	_, err = r.session.PutItemWithContext(ctx, input)
	log.Info(err)
	return err}