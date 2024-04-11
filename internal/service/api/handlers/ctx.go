package handlers

import (
	"context"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"

	"github.com/rarimo/humanornot-svc/internal/service/core"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	humanOrNotServiceCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxHumanornotSvc(humanOrNotSvc core.HumanornotSvc) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, humanOrNotServiceCtxKey, humanOrNotSvc)
	}
}

func HumanornotSvc(r *http.Request) core.HumanornotSvc {
	return r.Context().Value(humanOrNotServiceCtxKey).(core.HumanornotSvc).New()
}
