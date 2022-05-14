package web

import (
	"time"

	"github.com/google/uuid"
)

type RequestInfoKey struct{}

type RequestInfo struct {
	ID      uuid.UUID
	StartAt time.Time
}
