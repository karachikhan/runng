package migrations

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func MigratePlayerCreate() {

	svc := dynamodb.New(dynamodb.Options{
		Region:       "us-west-2",
		BaseEndpoint: aws.String("http://localhost:8000"),
		Credentials:  credentials.NewStaticCredentialsProvider("AKID", "SECRET", "SESSION"),
	})

	_, err := svc.CreateTable(context.Background(), &dynamodb.CreateTableInput{
		TableName: aws.String("Players"),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("PlayerID"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("PlayerID"),
				KeyType:       types.KeyTypeHash,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Table %s created successfully\n", "Players")
}
