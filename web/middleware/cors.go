package middleware

import (
	"context"
	"net/http"

	gocors "github.com/rs/cors"

	"github.com/bohdanch-w/go-api/entities/set"
	"github.com/bohdanch-w/go-api/web/api"
)

type CoorsMid struct {
	AllowAll         bool
	AllowedOrigins   set.StringSet
	AllowedMethods   set.StringSet
	AllowedHeaders   set.StringSet
	AllowCredentials bool
}

func (mid *CoorsMid) Wrap(h api.Handler) api.Handler {
	cors := gocors.New(gocors.Options{
		AllowedOrigins:   mid.Origins(),
		AllowedMethods:   mid.Methods(),
		AllowedHeaders:   mid.Headers(),
		AllowCredentials: mid.Credentials(),
	})

	f := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		var (
			err  error
			next = func(ww http.ResponseWriter, wr *http.Request) {
				err = h(ctx, ww, wr)
			}
		)

		cors.ServeHTTP(w, r, next)

		return err
	}

	return f
}

func (mid *CoorsMid) Origins() []string {
	if mid.AllowAll || mid.AllowedOrigins.IsEmpty() {
		return []string{"*"}
	}

	return mid.AllowedOrigins.Elements()
}

func (mid *CoorsMid) Methods() []string {
	if mid.AllowAll || mid.AllowedMethods.IsEmpty() {
		return []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodConnect,
			http.MethodTrace,
		}
	}

	return mid.AllowedMethods.Elements()
}

func (mid *CoorsMid) Headers() []string {
	if mid.AllowAll || mid.AllowedHeaders.IsEmpty() {
		return []string{"*"}
	}

	return mid.AllowedHeaders.Elements()
}

func (mid *CoorsMid) Credentials() bool {
	return mid.AllowAll || mid.AllowCredentials
}
