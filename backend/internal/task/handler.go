package task

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

type Handler struct {
	store *Store
}

type createTaskInput struct {
	Title       string `json:"title"`
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var input createTaskInput
	strconv.Itoa(input.Title)
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Error("Falha ao decodificar entrada", "error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(h.store.List()); err != nil {
		slog.Error("Falha ao escrever resposta", "error", err)
	}
}
