package repository

import (
	"api-users/models"
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const TableName = "user"

func CreateUser(ctx context.Context, user models.User) error {
	_, err := Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(TableName),
		Item: map[string]types.AttributeValue{
			"UserID":  &types.AttributeValueMemberS{Value: user.ID},
			"name":    &types.AttributeValueMemberS{Value: user.Name},
			"age":     &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", user.Age)},
			"rg":      &types.AttributeValueMemberS{Value: user.RG},
			"cpf":     &types.AttributeValueMemberS{Value: user.CPF},
			"address": &types.AttributeValueMemberS{Value: user.Address},
		},
	})
	return err
}

func UpdateUser(ctx context.Context, id string, user models.User) error {
	_, err := Client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(TableName),
		Key: map[string]types.AttributeValue{
			"UserID": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression: aws.String("SET #n = :name, age = :age, rg = :rg, cpf = :cpf, address = :address"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":name":    &types.AttributeValueMemberS{Value: user.Name},
			":age":     &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", user.Age)},
			":rg":      &types.AttributeValueMemberS{Value: user.RG},
			":cpf":     &types.AttributeValueMemberS{Value: user.CPF},
			":address": &types.AttributeValueMemberS{Value: user.Address},
		},
		ExpressionAttributeNames: map[string]string{
			"#n": "name",
		},
		ReturnValues: types.ReturnValueNone,
	})
	return err
}

func GetUser(ctx context.Context, id string) (*models.User, error) {
	result, err := Client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key: map[string]types.AttributeValue{
			"UserID": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		log.Println("Error searching item on DynamoDB:", err)
		return nil, err
	}
	if result.Item == nil {
		log.Println("User not found")
		return nil, nil
	}
	var user models.User
	err = attributevalue.UnmarshalMap(result.Item, &user)
	if err != nil {
		log.Println("conversion error", err)
		return nil, err
	}
	return &user, nil
}

func DeleteUser(ctx context.Context, id string) error {

	_, err := Client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(TableName),
		Key: map[string]types.AttributeValue{
			"UserID": &types.AttributeValueMemberS{Value: id},
		},
	})

	return err
}
