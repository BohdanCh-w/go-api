package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/bohdanch-w/go-api/web"
	"github.com/bohdanch-w/go-api/web/api"
)

type ErrorMid struct {
	Logger *log.Logger
}

func (mid *ErrorMid) Wrap(h api.Handler) api.Handler {
	f := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		info, ok := ctx.Value(web.RequestInfoKey{}).(*web.RequestInfo)
		if !ok {
			return h(ctx, w, r)
		}

		err := h(ctx, w, r)
		if err == nil {
			return err
		}

		webErr := err.(*web.WebError)

		mid.Logger.Printf("Request %s failed with status %d and error message %s",
			info.ID.String(), webErr.Status(), webErr.Error())

		if webErr.Status() > 0 {
			web.Abort(w, webErr)
		}

		return err
	}

	return f
}
