package migrations

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func MigrateTableCreate() {
	dbSvc := dynamodb.New(dynamodb.Options{
		Region:       "us-west-2",
		BaseEndpoint: aws.String("http://localhost:8000"),
		Credentials:  credentials.NewStaticCredentialsProvider("AKID", "SECRET", "SESSION"),
	})
	if err := createTable(context.Background(), dbSvc, "Table"); err != nil {
		panic(err)
	}
}

func createTable(ctx context.Context, svc *dynamodb.Client, tableName string) error {
	_, err := svc.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("TableID"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("TableID"),
				KeyType:       types.KeyTypeHash,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	})
	if err != nil {
		return err
	}

	fmt.Printf("Table %s created successfully\n", tableName)
	return nil
}
