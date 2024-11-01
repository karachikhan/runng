package api

import (
	"minhajuddinkhan/runng/runng/tables"
	"minhajuddinkhan/runng/store"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func NewJoinTableHandler(tableStore store.TableStore, playerStore store.PlayerStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tableID := mux.Vars(r)["table_id"]
		position := tablePosition(mux.Vars(r)["position"])

		table, err := tableStore.GetByID(r.Context(), uuid.MustParse(tableID))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		playerID := mux.Vars(r)["player_id"]
		player, err := playerStore.GetByID(r.Context(), uuid.MustParse(playerID))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tablePlayer, err := player.Join(table, position)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := tableStore.AddPlayerToTable(r.Context(), table.GetID(), tablePlayer); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func tablePosition(position string) tables.TablePlayerPosition {
	switch position {
	case "east":
		return tables.TablePlayerPositionEast
	case "west":
		return tables.TablePlayerPositionWest
	case "north":
		return tables.TablePlayerPositionNorth
	case "south":
		return tables.TablePlayerPositionSouth
	}
	return -1
}
