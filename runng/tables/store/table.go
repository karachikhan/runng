package store

type Table struct {
	ID      string
	Players []TablePlayer
}

type TablePlayer struct {
	ID       string
	Position string
}

type TableEventType int

const (
	EventTypeTableCreated TableEventType = iota
	EventTypePlayerJoined
)

type TableEvent struct {
	ID        int
	EventType TableEventType
	Table     Table
}
