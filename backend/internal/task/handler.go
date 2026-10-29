package task

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	store *Store
}

type createTaskInput struct {
	Title string `json:"title"`
}

type updateTaskInput struct {
	Title *string `json:"title"`
	Done  *bool   `json:"done"`
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

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	idText := r.PathValue("id")

	id, err := strconv.Atoi(idText)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	t, ok := h.store.Get(id)
	if !ok {
		http.Error(w, "id não encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		slog.Error("Falha ao escrever resposta", "error", err)
	}
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idText := r.PathValue("id")

	id, err := strconv.Atoi(idText)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	if !h.store.Delete(id) {
		http.Error(w, "tarefa não encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idText := r.PathValue("id")

	id, err := strconv.Atoi(idText)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	var input updateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Json inválido", http.StatusBadRequest)
		return
	}

	if input.Title == nil && input.Done == nil {
		http.Error(w, "informe title ou done", http.StatusBadRequest)
		return
	}

	if input.Title != nil {
		trimmed := strings.TrimSpace(*input.Title)
		if trimmed == "" {
			http.Error(w, "o título não pode ser vazio", http.StatusBadRequest)
			return
		}
		input.Title = &trimmed
	}

	t, ok := h.store.Update(id, input.Title, input.Done)
	if !ok {
		http.Error(w, "tarefa não encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		slog.Error("Falha ao escrever resposta", "error", err)
	}
}
