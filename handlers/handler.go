package handler

import (
	"github.com/google/uuid"
	"github.com/kanerix/gobyd/pkg/clock"
)

type Handler struct {
	clock clock.VClock
	pid   uuid.UUID
}

// Returns a new handler server.
func NewHandler() Handler {
	return Handler{clock.NewVClock(), uuid.New()}
}
