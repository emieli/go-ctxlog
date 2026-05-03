# Usage
```go
func main() {
	ctxLogFields := new(ctxlog.Fields)
	ctx := context.WithValue(context.Background(), ctxlog.Key, ctxLogFields)
	addCtxLogFields(ctx)

	slog.Info("success", "ctx_fields", ctxLogFields)
    // output: 2026/05/03 19:24:21 INFO success ctx_fields.herp=derp ctx_fields.flerp=smerp

func addCtxLogFields(ctx context.Context) {
	ctxlog, exist := ctx.Value(ctxlog.Key).(*ctxlog.Fields)
	if !exist {
		panic("ctxlog.Fields not found")
	}
	ctxlog.AddField(
		slog.String("herp", "derp"),
		slog.String("flerp", "smerp"),
	)
}
```

# Install
```bash
go get https://github.com/emieli/go-ctxlog
```
