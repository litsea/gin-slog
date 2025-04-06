package slog

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

var defaultLogExtraAttrs = map[string]string{
	"log.handler": "litsea.gin-slog",
}

type logConfig struct {
	extraAttrs map[string]string
}

func New(l *slog.Logger, opts ...Option) gin.HandlerFunc {
	cfg := &sloggin.Config{
		DefaultLevel:     slog.LevelInfo,
		ClientErrorLevel: slog.LevelWarn,
		ServerErrorLevel: slog.LevelError,

		WithUserAgent:      false,
		WithRequestID:      true,
		WithRequestBody:    false,
		WithRequestHeader:  false,
		WithResponseBody:   false,
		WithResponseHeader: false,
		WithSpanID:         false,
		WithTraceID:        false,

		Filters: []sloggin.Filter{},
	}

	lc := &logConfig{
		extraAttrs: defaultLogExtraAttrs,
	}

	for _, opt := range opts {
		opt(lc, cfg)
	}

	for k, v := range lc.extraAttrs {
		l = l.With(k, v)
	}

	return sloggin.NewWithConfig(l, *cfg)
}

func AddCustomAttributes(ctx *gin.Context, key string, value any) {
	sloggin.AddCustomAttributes(ctx, slog.Any(key, value))
}
