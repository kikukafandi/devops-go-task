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

// Handler untuk health-check API
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	log.Printf("INFO: Endpoint /health accessed from %s", r.RemoteAddr)

	json.NewEncoder(w).Encode(HealthResponse{
		Status:  "ok",
		Message: "Service is running gracefully",
	})
}

// Handler untuk halaman depan (HTML)
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Mencegah path lain selain "/" menampilkan halaman ini
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// HTML dan CSS Injeksi langsung
	html := `
	<!DOCTYPE html>
	<html lang="id">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Technical Task - DevOps</title>
		<style>
			body {
				font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
				background-color: #f8fafc;
				color: #334155;
				display: flex;
				justify-content: center;
				align-items: center;
				height: 100vh;
				margin: 0;
			}
			.card {
				background: white;
				padding: 40px;
				border-radius: 12px;
				box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1);
				text-align: center;
				max-width: 450px;
				width: 90%;
			}
			h1 { color: #0f172a; font-size: 24px; margin-bottom: 8px; }
			h2 { color: #64748b; font-size: 16px; font-weight: 500; margin-top: 0; margin-bottom: 24px; }
			.developer { font-size: 18px; font-weight: 600; color: #3b82f6; margin-bottom: 30px; }
			.btn {
				display: inline-block;
				padding: 12px 24px;
				color: white;
				background-color: #2563eb;
				text-decoration: none;
				border-radius: 6px;
				font-weight: 500;
				transition: background-color 0.2s;
			}
			.btn:hover { background-color: #1d4ed8; }
			.footer { margin-top: 32px; font-size: 12px; color: #94a3b8; }
		</style>
	</head>
	<body>
		<div class="card">
			<h1>Technical Task</h1>
			<h2>Fullstack & DevOps Engineer Position</h2>
			<div class="developer">Developed by Tirta Afandi</div>
			<a href="/health" class="btn">View API Health Check</a>
			<div class="footer">
				Automated Deployment via GitHub Actions & Docker<br>
				&copy; 2026
			</div>
		</div>
	</body>
	</html>
	`
	w.Write([]byte(html))
}

func main() {
	mux := http.NewServeMux()

	// Routing endpoint
	mux.HandleFunc("/", indexHandler)      // Route halaman utama (UI)
	mux.HandleFunc("/health", healthCheck) // Route API (JSON)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Println("INFO: Starting server on port :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ERROR: Server failed to start: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("INFO: Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("ERROR: Server forced to shutdown: %v\n", err)
	}

	log.Println("INFO: Server exiting gracefully")
}
