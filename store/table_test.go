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

func TestTableStore(t *testing.T) {
	dbSvc := dynamodb.New(dynamodb.Options{
		Region:       "us-west-2",
		BaseEndpoint: aws.String("http://localhost:8000"),
		Credentials:  credentials.NewStaticCredentialsProvider("AKID", "SECRET", "SESSION"),
	})

	tableStore := store.NewTableStore(dbSvc)
	theTable := tables.NewTable()
	err := tableStore.Create(context.Background(), theTable)
	assert.NoError(t, err)

	p := tables.NewPlayer()
	tp1, err := p.Join(theTable, tables.TablePlayerPositionEast)
	assert.NoError(t, err)

	p = tables.NewPlayer()
	tp2, err := p.Join(theTable, tables.TablePlayerPositionSouth)
	assert.NoError(t, err)
	// tp2,err:= tables.NewPlayer().

	err = tableStore.AddPlayerToTable(context.Background(), theTable.GetID(), tp1)
	assert.NoError(t, err)

	err = tableStore.AddPlayerToTable(context.Background(), theTable.GetID(), tp2)
	assert.NoError(t, err)

	x, err := tableStore.GetByID(context.Background(), theTable.GetID())
	assert.NoError(t, err)
	assert.Equal(t, theTable.GetID(), x.GetID())

	// err = tableStore.Delete(context.Background(), theTable.GetID())
	// assert.NoError(t, err)
}
