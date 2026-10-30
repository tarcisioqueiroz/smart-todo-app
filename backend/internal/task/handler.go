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

type erroResponse struct {
	Error string `json:"error"`
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	var input createTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Error("Falha ao decodificar entrada", "error", err)
		writeError(w, http.StatusBadRequest, "Json inválido")
		return
	}

	titleTrim := strings.TrimSpace(input.Title)
	if titleTrim == "" {
		writeError(w, http.StatusBadRequest, "Título da tarefa não pode ser vazio")
		return
	}

	t := h.store.Create(titleTrim)
	writeJSON(w, http.StatusCreated, t)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.store.List())
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id, ok := parseId(w, r)
	if !ok {
		return
	}

	t, ok := h.store.Get(id)
	if !ok {
		writeError(w, http.StatusNotFound, "tarefa não encontrada")
		return
	}

	writeJSON(w, http.StatusOK, t)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, ok := parseId(w, r)
	if !ok {
		return
	}

	if !h.store.Delete(id) {
		writeError(w, http.StatusNotFound, "tarefa não encontrada")
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, ok := parseId(w, r)
	if !ok {
		return
	}

	var input updateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "Json inválido")
		return
	}

	if input.Title == nil && input.Done == nil {
		writeError(w, http.StatusBadRequest, "informe title ou done")
		return
	}

	if input.Title != nil {
		trimmed := strings.TrimSpace(*input.Title)
		if trimmed == "" {
			writeError(w, http.StatusBadRequest, "o título não pode ser vazio")
			return
		}
		input.Title = &trimmed
	}

	t, ok := h.store.Update(id, input.Title, input.Done)
	if !ok {
		writeError(w, http.StatusNotFound, "tarefa não encontrada")
		return
	}

	writeJSON(w, http.StatusOK, t)
}

func parseId(w http.ResponseWriter, r *http.Request) (int, bool) {
	idText := r.PathValue("id")

	id, err := strconv.Atoi(idText)
	if err != nil {
		writeError(w, http.StatusBadRequest, "id inválido")
		return id, false
	}

	return id, true
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("Falha ao escrever resposta", "error", err)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, erroResponse{Error: message})
}
