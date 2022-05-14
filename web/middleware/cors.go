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
}

func (mid *CoorsMid) Wrap(h api.Handler) api.Handler {
	f := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		if mid.AllowAll {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		} else {
			if mid.AllowedOrigins.Contains(r.Host) {
				w.Header().Set("Access-Control-Allow-Origin", r.Host)
			}
		}

		return h(ctx, w, r)
	}

	return f
}
