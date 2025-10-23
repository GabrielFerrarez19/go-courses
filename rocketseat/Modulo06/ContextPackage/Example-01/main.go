package main

import (
	"context"
	"fmt"
)

type ctxKey string

func main() {
	doSomething(context.Background(), "rocketseat", "the best dev school")
}

func doSomething(ctx context.Context, name, desc string) {
	ctx = context.WithValue(ctx, ctxKey("schoolName"), name)
	ctx = context.WithValue(ctx, ctxKey("desc"), desc)
	doSomethingElse(ctx)
}

func doSomethingElse(ctx context.Context) {
	fmt.Printf("School %s is %s:\n", GetDesc(ctx, ctxKey("schoolName")), GetDesc(ctx, ctxKey("desc")))
}

func GetDesc(ctx context.Context, key ctxKey) string {
	return ctx.Value(key).(string)
}
