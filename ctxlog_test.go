package ctxlog_test

import (
	"context"
	"ctxlog"
	"log/slog"
	"testing"
)

func TestCtxLog(t *testing.T) {
	ctxLogFields := new(ctxlog.Fields)
	ctx := context.WithValue(context.Background(), ctxlog.Key, ctxLogFields)
	AddCtxLogFields(ctx)
	slog.Info("success", "ctxfields", ctxLogFields)
}

func AddCtxLogFields(ctx context.Context) {
	ctxlog, exist := ctx.Value(ctxlog.Key).(*ctxlog.Fields)
	if !exist {
		panic("ctxlog.Fields not found")
	}
	ctxlog.AddField(
		slog.String("herp", "derp"),
		slog.String("flerp", "smerp"),
	)
}
