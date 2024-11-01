package repository

import (
	"context"
	"errors"
	"minhajuddinkhan/runng/runng/tables"
	"minhajuddinkhan/runng/runng/tables/store"
)

type EventRepository interface {
	Load(context.Context, string) (tables.Table, error)
	Add(ctx context.Context, e store.TableEvent) error
}

func NewTableRepository() EventRepository {
	return &TableEventsRepository{
		events: make(map[string][]store.TableEvent),
	}
}

type TableEventsRepository struct {
	events map[string][]store.TableEvent
}

func (t *TableEventsRepository) Load(ctx context.Context, tableID string) (tables.Table, error) {
	events, ok := t.events[tableID]
	if !ok {
		return nil, errors.New("table not found")
	}
	lastEvent := events[len(events)-1]
	return tables.TableFromStore(lastEvent), nil
}

func (t *TableEventsRepository) Add(ctx context.Context, e store.TableEvent) error {
	if _, ok := t.events[e.Table.ID]; !ok {
		t.events[e.Table.ID] = make([]store.TableEvent, 0)
	}
	t.events[e.Table.ID] = append(t.events[e.Table.ID], e)
	return nil
}
