package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Database struct {
	connection *dynamodb.DynamoDB
	logMode    bool
}

type Interface interface {
	Health() bool
	ParseDynamoAtributeToStruct(entity interface{}, response map[string]*dynamodb.AttributeValue) error
	Find(condition map[string]interface{}, tableName string) (*dynamodb.GetItemOutput, error)
	Create(entity interface{}, tableName string) (*dynamodb.PutItemOutput, error)
	Update(condition map[string]interface{}, entity interface{}, tableName string) (*dynamodb.UpdateItemOutput, error)
	Delete(condition map[string]interface{}, tableName string) (*dynamodb.DeleteItemOutput, error)
}

func NewAdapter(con *dynamodb.DynamoDB) Interface {
	return &Database{
		connection: con,
		logMode:    false,
	}
}

func (db *Database) Health() bool {
	_, err := db.connection.ListTables(&dynamodb.ListTablesInput{})
	return err == nil
}

func (db *Database) ParseDynamoAtributeToStruct(entity interface{}, response map[string]*dynamodb.AttributeValue) error {
	return dynamodbattribute.UnmarshalMap(response, entity)
}

func (db *Database) Find(condition map[string]interface{}, tableName string) (*dynamodb.GetItemOutput, error) {
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: conditionParsed,
	}
	return db.connection.GetItem(input)
}

func (db *Database) Create(entity interface{}, tableName string) (*dynamodb.PutItemOutput, error) {
	entityParsed, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(tableName),
	}
	return db.connection.PutItem(input)
}

func (db *Database) Update(condition map[string]interface{}, entity interface{}, tableName string) (*dynamodb.UpdateItemOutput, error) {
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		return nil, err
	}
	entityParsed, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: entityParsed,
		TableName: aws.String(tableName),
		Key: conditionParsed,
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Rating = :r"),
	}
	return db.connection.UpdateItem(input)
}

func (db *Database) Delete(condition map[string]interface{}, tableName string) (*dynamodb.DeleteItemOutput, error) {
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.DeleteItemInput{
		Key: conditionParsed,
		TableName: aws.String(tableName),
	}
	return db.connection.DeleteItem(input)
}
