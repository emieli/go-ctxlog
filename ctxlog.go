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

type ctxKey string

var Key = ctxKey("ctxlog")

type Fields struct {
	mtx    sync.Mutex
	fields []slog.Attr
}

func (f *Fields) AddField(attr ...slog.Attr) {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	f.fields = append(f.fields, attr...)
}

// Implements log.LogValuer interface
func (f *Fields) LogValue() slog.Value {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	return slog.GroupValue(f.fields...)
}

// NewGroup creates a new Group that allows further nesting/grouping of slog attrs.
// This is useful if you for example want all logs from each function to have
// its own group in the final log. Create a new group and put all your logs in the group.
// Don't forget to add the logs to the 'clog' parent via the AppendTo funcion.
func (f Fields) NewGroup(name string) Group {
	return Group{name: name}
}

// Group allows further nesting/grouping of slog attrs.
// This is useful if you for example want all logs from each function to have
// its own group in the final log.
type Group struct {
	name  string
	attrs []slog.Attr
}

func (g *Group) AddField(attr ...slog.Attr) {
	g.attrs = append(g.attrs, attr...)
}

// AppendTo appends all logs collected by the group to the 'clog' parent.
// You can defer this function call, ensuring we have received all logs before sending to parent.
func (g *Group) AppendTo(clog *Fields) {
	clog.AddField(slog.GroupAttrs(g.name, g.attrs...))
}
