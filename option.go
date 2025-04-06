package slog

import (
	"log/slog"

	sloggin "github.com/samber/slog-gin"
)

type Option func(_ *logConfig, cfg *sloggin.Config)

func WithDefaultLevel(lv slog.Level) Option {
	return func(_ *logConfig, cfg *sloggin.Config) {
		cfg.DefaultLevel = lv
	}
}

func WithClientErrorLevel(lv slog.Level) Option {
	return func(_ *logConfig, cfg *sloggin.Config) {
		cfg.ClientErrorLevel = lv
	}
}

func WithServerErrorLevel(lv slog.Level) Option {
	return func(_ *logConfig, cfg *sloggin.Config) {
		cfg.ServerErrorLevel = lv
	}
}

func WithRequestID(v bool) Option {
	return func(_ *logConfig, cfg *sloggin.Config) {
		cfg.WithRequestID = v
	}
}

func WithUserAgent(v bool) Option {
	return func(_ *logConfig, cfg *sloggin.Config) {
		cfg.WithUserAgent = v
	}
}

func WithRequestHeader(v bool) Option {
	return func(_ *logConfig, cfg *sloggin.Config) {
		cfg.WithRequestHeader = v
	}
}

func WithRequestBody(v bool) Option {
	return func(_ *logConfig, cfg *sloggin.Config) {
		cfg.WithRequestBody = v
	}
}

func WithResponseHeader(v bool) Option {
	return func(_ *logConfig, cfg *sloggin.Config) {
		cfg.WithResponseHeader = v
	}
}

func WithResponseBody(v bool) Option {
	return func(_ *logConfig, cfg *sloggin.Config) {
		cfg.WithResponseBody = v
	}
}

func WithFilters(fs ...sloggin.Filter) Option {
	return func(_ *logConfig, cfg *sloggin.Config) {
		if len(fs) > 0 {
			cfg.Filters = fs
		}
	}
}

func WithSpanID(v bool) Option {
	return func(_ *logConfig, cfg *sloggin.Config) {
		cfg.WithSpanID = v
	}
}

func WithTraceID(v bool) Option {
	return func(_ *logConfig, cfg *sloggin.Config) {
		cfg.WithTraceID = v
	}
}

func WithExtraAttrs(attrs map[string]string) Option {
	return func(lc *logConfig, _ *sloggin.Config) {
		lc.extraAttrs = attrs
	}
}
