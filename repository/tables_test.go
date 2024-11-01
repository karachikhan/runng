package repository_test

import (
	"context"
	"minhajuddinkhan/runng/repository"
	"minhajuddinkhan/runng/runng/tables"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventRepository_Load(t *testing.T) {

	tablesRepository := repository.NewTableRepository()
	tablesService := tables.NewTableService(tablesRepository)
	table, err := tablesRepository.Load(context.Background(), "1")
	assert.Error(t, err)
	assert.Nil(t, table)

	table, err = tablesService.CreateTable(context.Background())
	assert.NoError(t, err)
	tableID := table.GetID().String()
	assert.NoError(t, err)

	table, err = tablesRepository.Load(context.Background(), table.GetID().String())
	assert.NoError(t, err)
	assert.NotNil(t, table)
	assert.Equal(t, tableID, table.GetID().String())
}
