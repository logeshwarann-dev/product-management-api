package logger

import (
	"context"
	"log/slog"
	"os"
)

type CtxKey struct{}

func AddLoggerToCtx(ctx context.Context, slogger *slog.Logger) context.Context {
	if slogger == nil {
		return ctx
	}
	if loggerWithCtx, ok := ctx.Value(CtxKey{}).(*slog.Logger); ok && loggerWithCtx == slogger {
		return ctx
	}
	return context.WithValue(ctx, CtxKey{}, slogger)
}

func GetLoggerUsingCtx(ctx context.Context) *slog.Logger {
	if slogger, ok := ctx.Value(CtxKey{}).(*slog.Logger); ok {
		return slogger
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
}
