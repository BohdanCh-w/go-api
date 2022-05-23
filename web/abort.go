package web

import (
	"net/http"
)

func Abort(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	if err, ok := err.(*WebError); ok {
		w.WriteHeader(err.Code)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write([]byte(`{"error": "` + err.Error() + `"}`))
}
