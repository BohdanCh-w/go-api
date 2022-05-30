package middleware

import (
	"context"
	"net/http"

	"github.com/bohdanch-w/go-api/entities/set"
	"github.com/bohdanch-w/go-api/web/api"
)

type CoorsMid struct {
	AllowAll       bool
	AllowedOrigins set.StringSet
	AllowedMethods set.StringSet
}

func (mid *CoorsMid) Wrap(h api.Handler) api.Handler {
	f := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		if mid.AllowAll {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		} else {
			if mid.AllowedOrigins.Contains(r.Host) {
				w.Header().Set("Access-Control-Allow-Origin", r.Host)
			}

			if mid.AllowedMethods.Contains(r.Method) {
				w.Header().Set("Access-Control-Allow-Methods", r.Method)
			}
		}

		return h(ctx, w, r)
	}

	return f
}
