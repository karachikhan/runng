package api

import (
	"encoding/json"
	"minhajuddinkhan/runng/runng/tables"
	"minhajuddinkhan/runng/store"
	"net/http"
)

func NewCreatePlayerHandler(playerStore store.PlayerStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		player := tables.NewPlayer()
		if err := playerStore.CreatePlayer(r.Context(), &player); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(PlayerCreateResponse{
			PlayerID: player.GetID().String(),
		})
		w.WriteHeader(http.StatusCreated)
	})
}

type PlayerCreateResponse struct {
	PlayerID string `json:"player_id"`
}
