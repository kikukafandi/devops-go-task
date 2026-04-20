package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Response struct untuk standar output JSON
type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Handler untuk health-check
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Logging dasar setiap kali endpoint diakses
	log.Printf("INFO: Endpoint /health accessed from %s", r.RemoteAddr)

	json.NewEncoder(w).Encode(HealthResponse{
		Status:  "ok",
		Message: "Service is running gracefully",
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthCheck)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Menjalankan server dalam goroutine agar tidak memblokir eksekusi selanjutnya
	go func() {
		log.Println("INFO: Starting server on port :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ERROR: Server failed to start: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Menunggu sinyal masuk
	<-quit
	log.Println("INFO: Shutting down server...")

	// Memberikan waktu maksimal 5 detik untuk menyelesaikan request yang masih berjalan
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("ERROR: Server forced to shutdown: %v\n", err)
	}

	log.Println("INFO: Server exiting")
}
