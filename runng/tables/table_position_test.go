package tables_test

import (
	"minhajuddinkhan/runng/runng/tables"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTablePosition(t *testing.T) {
	pos := tables.TablePlayerPositionEast
	assert.Equal(t, tables.TablePlayerPositionNorth, pos.Next())
	pos = pos.Next()
	assert.Equal(t, tables.TablePlayerPositionWest, pos.Next())
	pos = pos.Next()
	assert.Equal(t, tables.TablePlayerPositionSouth, pos.Next())
	pos = pos.Next()
	assert.Equal(t, tables.TablePlayerPositionEast, pos.Next())

}
