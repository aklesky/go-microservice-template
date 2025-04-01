package main

import (
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("Starting server on :8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received request", "method", r.Method, "url", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Server failed", "error", err)
	}
}
