package rest

import (
	"sync"

	"github.com/google/uuid"
	"github.com/kanerix/gobyd/pkg/clock"
)

type Handler struct {
	sync.RWMutex
	clock  clock.VClock
	nodeID uuid.UUID
	queue  []uuid.UUID
	peers  []string
}

// Returns a new handler for the REST API.
func NewHandler(peers []string) Handler {
	return Handler{
		RWMutex: sync.RWMutex{},
		clock:   clock.VClock{},
		nodeID:  uuid.New(),
		queue:   make([]uuid.UUID, 0),
		peers:   peers,
	}
}
