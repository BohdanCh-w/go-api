package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/bohdanch-w/go-api/web"
	"github.com/bohdanch-w/go-api/web/api"
)

type IdentityMid struct {
	Logger *log.Logger
}

func (mid *IdentityMid) Wrap(h api.Handler) api.Handler {
	f := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		var (
			id    = uuid.New()
			start = time.Now()
		)

		ctx = context.WithValue(ctx, web.RequestInfoKey{}, &web.RequestInfo{
			ID:      id,
			StartAt: start,
		})

		mid.Logger.Printf("Request %s started at %s:\n\t method: %s url: %s",
			id.String(), start.Format("02-Jan-2006 15:04:05.999"), r.Method, r.URL.String())

		return h(ctx, w, r)
	}

	return f
}
