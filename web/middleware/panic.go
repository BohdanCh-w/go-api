package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/google/uuid"
	"github.com/bohdanch-w/go-api/web"
	"github.com/bohdanch-w/go-api/web/api"
)

type PanicMid struct {
	Logger *log.Logger
}

func (mid *PanicMid) Wrap(h api.Handler) api.Handler {
	f := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		info, ok := ctx.Value(web.RequestInfoKey{}).(*web.RequestInfo)
		if !ok {
			info = &web.RequestInfo{ID: uuid.Nil}
		}

		defer func() {
			if r := recover(); r != nil {
				stack := debug.Stack()

				mid.Logger.Printf("Request %s got fatal server error %q on %#v", info.ID.String(), stack, r)

				web.Abort(w, &web.WebError{
					Code: http.StatusInternalServerError,
					Err:  fmt.Errorf("Fatal server error"),
				})
			}
		}()

		return h(ctx, w, r)
	}

	return f
}
