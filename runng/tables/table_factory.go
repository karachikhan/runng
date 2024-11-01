package tables

import (
	"minhajuddinkhan/runng/runng/cards"
	"minhajuddinkhan/runng/runng/tables/store"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

func TableFromStore(e store.TableEvent) Table {
	players := make(map[TablePlayerPosition]*TablePlayer)
	for _, p := range e.Table.Players {
		pos := positionFromString(p.Position)
		players[pos] = &TablePlayer{
			tableID:  uuid.MustParse(e.Table.ID),
			position: pos,
			Player: &Player{
				id: uuid.MustParse(p.ID),
			},
			cards: make([]cards.Card, 0),
		}
	}
	return &table{
		id:      uuid.MustParse(e.Table.ID),
		players: players,
	}
}

type DynamoFactory struct{}

func (df *DynamoFactory) TableFromDynamoAttributes(result *dynamodb.GetItemOutput) Table {
	res := &table{
		id:      uuid.MustParse(result.Item["TableID"].(*types.AttributeValueMemberS).Value),
		players: make(map[TablePlayerPosition]*TablePlayer),
	}
	players := result.Item["Players"].(*types.AttributeValueMemberL).Value
	for _, v := range players {
		pos := positionFromString(v.(*types.AttributeValueMemberM).Value["Position"].(*types.AttributeValueMemberS).Value)
		p := &TablePlayer{
			tableID:  res.id,
			position: pos,
			cards:    make([]cards.Card, 0),
			Player: &Player{
				id: uuid.MustParse(v.(*types.AttributeValueMemberM).Value["PlayerID"].(*types.AttributeValueMemberS).Value),
			},
		}
		res.players[pos] = p
	}
	return res
}

func (df *DynamoFactory) PlayerFromDynamoAttributes(result *dynamodb.GetItemOutput) *Player {
	return &Player{
		id: uuid.MustParse(result.Item["PlayerID"].(*types.AttributeValueMemberS).Value),
	}
}

func positionFromString(pos string) TablePlayerPosition {
	switch pos {
	case "East":
		return TablePlayerPositionEast
	case "South":
		return TablePlayerPositionSouth
	case "West":
		return TablePlayerPositionWest
	case "North":
		return TablePlayerPositionNorth
	}
	return -1
}
