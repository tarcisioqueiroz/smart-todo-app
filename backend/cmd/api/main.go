package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/tarcisioqueiroz/smart-todo-app/backend/internal/task"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /tasks", listTasksHandler)

	port := getEnv("PORT", "8080")
	addr := ":" + port

	slog.Info("Servidor escutando na porta", "addr", addr)
	err := http.ListenAndServe(addr, mux)
	slog.Error("não foi possível iniciar o servidor", "error", err)
	os.Exit(1)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func listTasksHandler(w http.ResponseWriter, r *http.Request) {
	list := []task.Task{
		{ID: 1, Title: "estudar go", CreatedAt: time.Now()},
		{ID: 2, Title: "revisar PR9", CreatedAt: time.Now()},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(list); err != nil {
		slog.Error("Falha de ao escrever resposta", "error", err)
	}
}

func getEnv(key, fallback string) string {
	valor := os.Getenv(key)
	if valor != "" {
		return valor
	}

	return fallback
}
