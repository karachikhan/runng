package tables_test

import (
	"context"
	"minhajuddinkhan/runng/repository"
	"minhajuddinkhan/runng/runng/tables"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableService(t *testing.T) {
	tableRepository := repository.NewTableRepository()
	ts := tables.NewTableService(tableRepository)

	// Monday
	theTable, err := ts.CreateTable(context.Background())
	assert.NoError(t, err)
	tableID := theTable.GetID().String()

	// Tuesday
	table, err := tableRepository.Load(context.Background(), tableID)
	assert.NoError(t, err)
	assert.NotNil(t, table)

	// Wednesday
	err = ts.JoinPlayer(context.Background(), tableID, tables.NewPlayer(), tables.TablePlayerPositionNorth)
	assert.NoError(t, err)
	err = ts.JoinPlayer(context.Background(), tableID, tables.NewPlayer(), tables.TablePlayerPositionEast)
	assert.NoError(t, err)
	err = ts.JoinPlayer(context.Background(), tableID, tables.NewPlayer(), tables.TablePlayerPositionSouth)
	assert.NoError(t, err)
	err = ts.JoinPlayer(context.Background(), tableID, tables.NewPlayer(), tables.TablePlayerPositionWest)
	assert.NoError(t, err)

	// Thursday
	err = ts.JoinPlayer(context.Background(), tableID, tables.NewPlayer(), tables.TablePlayerPositionWest)
	assert.Error(t, err)
}
