// package ctxlog allow passing logging fields to a context.
// This allow passing more information up the call stack, allowing the
// function at the top to create one log entry with information provided
// by called funcions.
// This provide more helpful logs as one log line contain all info,
// instead of each function creating its own log line with only its own info.
//
//	https://rednafi.com/shards/2026/04/no-stacked-loglines/
package ctxlog

import (
	"log/slog"
	"sync"
)

type Fields struct {
	mtx    sync.Mutex
	fields []slog.Attr
}

func (f *Fields) AddField(attr ...slog.Attr) {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	f.fields = append(f.fields, attr...)
}

func (f *Fields) LogValue() slog.Value {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	return slog.GroupValue(f.fields...)
}

type ctxKey string

var Key = ctxKey("ctxlog")
