package web

import (
	"net/http"
)

func Abort(w http.ResponseWriter, weberr *WebError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(weberr.Code)

	w.Write([]byte(`{"error": "` + weberr.Error() + `"}`))
}
