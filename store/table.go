package store

import (
	"context"
	"minhajuddinkhan/runng/runng/tables"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type TableStore interface {
	Create(ctx context.Context, t tables.Table) error
	Delete(ctx context.Context, tableID uuid.UUID) error
	GetByID(ctx context.Context, tableID uuid.UUID) (tables.Table, error)
	AddPlayerToTable(ctx context.Context, tableID uuid.UUID, p *tables.TablePlayer) error
}

type tableStore struct {
	client  *dynamodb.Client
	factory *tables.DynamoFactory
}

func NewTableStore(c *dynamodb.Client) TableStore {
	return &tableStore{
		client:  c,
		factory: &tables.DynamoFactory{},
	}
}

func (ts *tableStore) Create(ctx context.Context, t tables.Table) error {
	_, err := ts.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("Table"),
		Item: map[string]types.AttributeValue{
			"TableID": &types.AttributeValueMemberS{Value: t.GetID().String()},
		},
	})
	return err
}

func (ts *tableStore) Delete(ctx context.Context, tableID uuid.UUID) error {
	_, err := ts.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String("Table"),
		Key: map[string]types.AttributeValue{
			"TableID": &types.AttributeValueMemberS{Value: tableID.String()},
		},
	})
	return err
}

func (ts *tableStore) GetByID(ctx context.Context, tableID uuid.UUID) (tables.Table, error) {
	resp, err := ts.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String("Table"),
		Key: map[string]types.AttributeValue{
			"TableID": &types.AttributeValueMemberS{Value: tableID.String()},
		},
	})
	if err != nil {
		return nil, err
	}

	return ts.factory.TableFromDynamoAttributes(resp), nil
}

func (ts *tableStore) AddPlayerToTable(ctx context.Context, tableID uuid.UUID, p *tables.TablePlayer) error {
	_, err := ts.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		Key: map[string]types.AttributeValue{
			"TableID": &types.AttributeValueMemberS{Value: tableID.String()},
		},
		TableName: aws.String("Table"),
		AttributeUpdates: map[string]types.AttributeValueUpdate{
			"Players": {
				Action: types.AttributeActionAdd,
				Value: &types.AttributeValueMemberL{
					Value: []types.AttributeValue{
						&types.AttributeValueMemberM{
							Value: map[string]types.AttributeValue{
								"PlayerID": &types.AttributeValueMemberS{Value: p.GetID().String()},
								"Position": &types.AttributeValueMemberS{Value: p.GetPosition().String()},
							},
						},
					},
				},
			},
		},
	})
	return err
}
