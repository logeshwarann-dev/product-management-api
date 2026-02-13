package middleware

import (
	"go-native-http/internal/logger"
	"log/slog"
	"net/http"
)

func InjectLogger(slogger *slog.Logger, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxWithLogger := logger.AddLoggerToCtx(r.Context(), slogger)
		r = r.Clone(ctxWithLogger)
		next.ServeHTTP(w, r)
	}
}

func AddLoggerToRequest(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log := logger.GetLoggerUsingCtx(r.Context())
		log.Info("request", "path", r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
