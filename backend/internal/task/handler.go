package task

import (
	"log/slog"
	"net/http"
	"encoding/json"
)

type Handler struct {
	store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(h.store.List()); err != nil {
		slog.Error("Falha ao escrever resposta", "error", err)
	}
}
