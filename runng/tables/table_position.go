package tables

type TablePlayerPosition int

const (
	TablePlayerPositionEast TablePlayerPosition = iota
	TablePlayerPositionSouth
	TablePlayerPositionWest
	TablePlayerPositionNorth
)

func (p TablePlayerPosition) String() string {
	switch p {
	case TablePlayerPositionEast:
		return "East"
	case TablePlayerPositionSouth:
		return "South"
	case TablePlayerPositionWest:
		return "West"
	case TablePlayerPositionNorth:
		return "North"
	}
	return ""
}

func (pos TablePlayerPosition) Next() TablePlayerPosition {
	switch pos {
	case TablePlayerPositionEast:
		return TablePlayerPositionNorth
	case TablePlayerPositionNorth:
		return TablePlayerPositionWest
	case TablePlayerPositionWest:
		return TablePlayerPositionSouth
	case TablePlayerPositionSouth:
		return TablePlayerPositionEast
	}
	return TablePlayerPositionEast
}

func (pos TablePlayerPosition) Prev() TablePlayerPosition {
	switch pos {
	case TablePlayerPositionEast:
		return TablePlayerPositionSouth
	case TablePlayerPositionSouth:
		return TablePlayerPositionWest
	case TablePlayerPositionWest:
		return TablePlayerPositionNorth
	case TablePlayerPositionNorth:
		return TablePlayerPositionEast
	}
	return TablePlayerPositionEast
}
