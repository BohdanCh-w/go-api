package web

import (
	"net/http"
)

func Abort(w http.ResponseWriter, err error) error {
	w.Header().Set("Content-Type", "application/json")

	if err, ok := err.(*WebError); ok {
		w.WriteHeader(err.Code)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	_, writeErr := w.Write([]byte(`{"error": "` + err.Error() + `"}`))

	return writeErr
}
