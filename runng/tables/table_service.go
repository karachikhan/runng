package tables

import (
	"context"
	"minhajuddinkhan/runng/runng/tables/store"
	"time"

	"github.com/google/uuid"
)

func NewTableService(eventStore EventStore) *TableService {
	return &TableService{
		eventStore: eventStore,
	}
}

type EventStore interface {
	Load(context.Context, string) (Table, error)
	Add(ctx context.Context, e store.TableEvent) error
}

type TableService struct {
	eventStore EventStore
}

func (sf *TableService) CreateTable(ctx context.Context) (Table, error) {
	t := &table{
		id:      uuid.New(),
		players: make(map[TablePlayerPosition]*TablePlayer),
	}
	if err := sf.eventStore.Add(ctx, store.TableEvent{
		ID:        int(time.Now().Unix()),
		EventType: store.EventTypeTableCreated,
		Table:     store.Table{ID: t.id.String()},
	}); err != nil {
		return nil, err
	}

	return t, nil
}

func (sf *TableService) JoinPlayer(ctx context.Context, tableID string, p Player, pos TablePlayerPosition) error {
	t, err := sf.eventStore.Load(ctx, tableID)
	if err != nil {
		return err
	}
	tablePlayer, err := t.Join(&p, pos)
	if err != nil {
		return err
	}

	theTable := t.(*table)
	players := make([]store.TablePlayer, 0)
	for _, v := range theTable.players {
		players = append(players, store.TablePlayer{
			ID:       v.Player.GetID().String(),
			Position: v.position.String(),
		})
	}

	players = append(players, store.TablePlayer{
		ID:       tablePlayer.GetID().String(),
		Position: tablePlayer.position.String(),
	})

	return sf.eventStore.Add(ctx, store.TableEvent{
		ID:        int(time.Now().Unix()),
		EventType: store.EventTypePlayerJoined,
		Table: store.Table{
			ID:      t.GetID().String(),
			Players: players,
		},
	})
}
