package store_test

import (
	"context"
	"minhajuddinkhan/runng/runng/tables"
	"minhajuddinkhan/runng/store"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/assert"
)

func TestPlayerStore(t *testing.T) {
	dbSvc := dynamodb.New(dynamodb.Options{
		Region:       "us-west-2",
		BaseEndpoint: aws.String("http://localhost:8000"),
		Credentials:  credentials.NewStaticCredentialsProvider("AKID", "SECRET", "SESSION"),
	})

	playerStore := store.NewPlayerStore(dbSvc)
	player := tables.NewPlayer()
	err := playerStore.CreatePlayer(context.Background(), &player)
	assert.NoError(t, err)

	p, err := playerStore.GetByID(context.Background(), player.GetID())
	assert.NoError(t, err)
	assert.Equal(t, player.GetID(), p.GetID())
}
