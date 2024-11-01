package api

import (
	"encoding/json"
	"minhajuddinkhan/runng/runng/tables"
	"minhajuddinkhan/runng/store"
	"net/http"
)

func NewCreateTableHandler(store store.TableStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		table := tables.NewTable()
		if err := store.Create(r.Context(), table); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(TableCreateResponse{
			TableID: table.GetID().String(),
		})
		w.WriteHeader(http.StatusCreated)
	})
}

type TableCreateResponse struct {
	TableID string `json:"table_id"`
}
