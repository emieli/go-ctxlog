package ctxlog_test

import (
	"context"
	"log/slog"
	"os"
	"ctxlog"
	"testing"
	"time"
)

func TestCtxLog(t *testing.T) {
	clog := new(ctxlog.Fields)
	ctx := context.WithValue(t.Context(), ctxlog.Key, clog)

	// Populate clog with slog.Attrs
	addCtxLogFields(ctx)

	// Print logs from context.
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	log.Info("print", "ctxlog", clog)

	// Example output:
	// {
	//   "time":"2026-10-09T13:31:08.2Z",
	//   "level":"INFO",
	//   "msg":"print",
	//   "ctxlog":{"func_addCtxLogFields":{"herp":"derp","herp":"derp2","answer_to_everything":42,"time_taken_ms":100}},
	// }
}

func addCtxLogFields(ctx context.Context) {
	clog := ctx.Value(ctxlog.Key).(*ctxlog.Fields) // Will panic if not found
	
	start := time.Now()
	time.Sleep(100 * time.Millisecond)

	clog.AddField(slog.Group("func addCtxLogFields",
		// Duplicate keys are allowed. Note that JSON doesn't allow multiple fields.
		// Tools like "jq" will only keep the last occurence, only showing the "derp2" value
		slog.String("herp", "derp"),
		slog.String("herp", "derp2"),

		slog.Int("answer to everything", 42),
		slog.Duration("time taken (ms)", time.Duration(time.Since(start).Milliseconds())),
	))
}
