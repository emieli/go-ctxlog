package ctxlog_test

import (
	"context"
	"ctxlog" // You have to fix this import, it's probably in a different subdirectory in your code
	"log/slog"
	"os"
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
	log.Info("print", "clog", clog)

	// Example output:
	// {
	//     "time": "2026-11-01T16:44:00.728694992+01:00",
	//     "level": "INFO",
	//     "msg": "print",
	//     "ctxlog": {
	//         "addCtxLogFields": {
	//             "herp": "derp2",
	//             "answer to everything": 42,
	//             "time taken (ms)": 100
	//         },
	//         "nested1": {
	//             "test1": "test1",
	//             "test2": "test2"
	//         },
	//         "nested2": {
	//             "test1": "test1",
	//             "test2": "test2"
	//         }
	//     }
	// }
}

func addCtxLogFields(ctx context.Context) {
	clog := ctx.Value(ctxlog.Key).(*ctxlog.Fields) // Will panic if not found

	start := time.Now()
	time.Sleep(100 * time.Millisecond)

	clog.AddField(slog.Group("addCtxLogFields",
		// Duplicate keys are allowed. Note that JSON doesn't allow multiple fields.
		// Tools like "jq" will only keep the last occurence, only showing the "derp2" value
		slog.String("herp", "derp"),
		slog.String("herp", "derp2"),

		slog.Int("answer to everything", 42),
		slog.Duration("time taken (ms)", time.Duration(time.Since(start).Milliseconds())),
	))

	// Allow nested function to return its logs in their own sections/groups:
	nested1(ctx)
	nested2(ctx)
}

// nested1 adds the test1/test2 logs in its own 'nested1' group
func nested1(ctx context.Context) {
	clog := ctx.Value(ctxlog.Key).(*ctxlog.Fields)
	cGroup := clog.NewGroup("nested1")
	defer cGroup.AppendTo(clog)

	cGroup.AddField(slog.String("test1", "test1"))
	cGroup.AddField(slog.String("test2", "test2"))
}

// nested2 adds the test1/test2 logs in its own 'nested2' group
func nested2(ctx context.Context) {
	clog := ctx.Value(ctxlog.Key).(*ctxlog.Fields)
	cGroup := clog.NewGroup("nested2")
	defer cGroup.AppendTo(clog)

	cGroup.AddField(slog.String("test1", "test1"), slog.String("test2", "test2"))
}
