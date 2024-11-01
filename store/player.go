package store

import (
	"context"
	"minhajuddinkhan/runng/runng/tables"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

// PlayerStore

type PlayerStore interface {
	CreatePlayer(ctx context.Context, p *tables.Player) error
	GetByID(ctx context.Context, playerID uuid.UUID) (*tables.Player, error)
}

type playerStore struct {
	client  *dynamodb.Client
	factory *tables.DynamoFactory
}

func NewPlayerStore(c *dynamodb.Client) PlayerStore {
	return &playerStore{
		client:  c,
		factory: &tables.DynamoFactory{},
	}
}

func (ps *playerStore) CreatePlayer(ctx context.Context, p *tables.Player) error {
	_, err := ps.client.PutItem(ctx, &dynamodb.PutItemInput{
		Item: map[string]types.AttributeValue{
			"PlayerID": &types.AttributeValueMemberS{Value: p.GetID().String()},
		},
		TableName: aws.String("Players"),
	})
	return err
}

func (ps *playerStore) GetByID(ctx context.Context, playerID uuid.UUID) (*tables.Player, error) {
	resp, err := ps.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String("Players"),
		Key: map[string]types.AttributeValue{
			"PlayerID": &types.AttributeValueMemberS{Value: playerID.String()},
		},
	})
	if err != nil {
		return nil, err
	}
	return ps.factory.PlayerFromDynamoAttributes(resp), nil
}
