package logging

import (
	"context"

	"go.uber.org/zap"
)

type logCtxKey struct{}

func Ctx(ctx context.Context) (logger *zap.SugaredLogger) {
	logger, ok := ctx.Value(logCtxKey{}).(*zap.SugaredLogger)
	if !ok {
		logger = zap.S()
	}
	return
}

func WithCtx(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, logCtxKey{}, logger)
}
