package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Respond(w http.ResponseWriter, status int, v interface{}) error {
	w.WriteHeader(status)
	if v == nil {
		return nil
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(v); err != nil {
		return NewError(-1, fmt.Errorf("web: write data failed: %w", err))
	}

	return nil
}
