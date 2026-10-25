package task

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
)

type Handler struct {
	store *Store
}

type createTaskInput struct {
	Title string `json:"title"`
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	var input createTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Error("Falha ao decodificar entrada", "error", err)
		http.Error(w, "Json inválido", http.StatusBadRequest)
		return
	}

	titleTrim := strings.TrimSpace(input.Title)
	if titleTrim == "" {
		http.Error(w, "Título da tarefa não pode ser vazio", http.StatusBadRequest)
		return
	}

	t := h.store.Create(titleTrim)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		slog.Error("Falha ao escrever resposta", "error", err)
	}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(h.store.List()); err != nil {
		slog.Error("Falha ao escrever resposta", "error", err)
	}
}
