package repository

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func CreateTable(client *dynamodb.Client) {
	tableName := "user"

	_, err := client.DescribeTable(context.TODO(), &dynamodb.DescribeTableInput{
		TableName: &tableName,
	})

	if err == nil {
		fmt.Println("Table exists:", tableName)
		return
	}

	input := &dynamodb.CreateTableInput{
		TableName: &tableName,
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String("UserID"), AttributeType: types.ScalarAttributeTypeS},
		},
		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String("UserID"), KeyType: types.KeyTypeHash},
		},
		BillingMode: types.BillingModePayPerRequest,
	}
	_, err = client.CreateTable(context.TODO(), input)
	if err != nil {
		fmt.Println("Error create table:", err)
	} else {
		fmt.Println("Table create sucessfuly:", tableName)
	}
}
