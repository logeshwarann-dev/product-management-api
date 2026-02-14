package main

import (
	"fmt"
	"go-native-http/internal/middleware"
	"go-native-http/internal/router"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const serverPort = 8080

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	r := router.Router{
		Mux: http.NewServeMux(),
	}
	router := r.New(nil)
	handler := r.AddCORS(router)
	wrappedRouter := middleware.InjectLogger(log, middleware.AddLoggerToRequest(handler))
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", serverPort),
		ReadHeaderTimeout: 5 * time.Second,
		Handler:           wrappedRouter,
	}

	log.Info("starting the server on port %d", serverPort)
	if err := server.ListenAndServe(); err != nil {
		log.Error("error running server", "error", err)
	}

}
