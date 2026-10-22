package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)

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

func getEnv(key, fallback string) string {
	valor := os.Getenv(key)
	if valor != "" {
		return valor
	}

	return fallback
}
