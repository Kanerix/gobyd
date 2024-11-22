package muex

import (
	"sync"

	"github.com/google/uuid"
	"github.com/kanerix/gobyd/pkg/clock"
)

type Handler struct {
	sync.Mutex
	clock  clock.VClock
	nodeID uuid.UUID
	queue  []uuid.UUID
}

func NewHandler() Handler {
	return Handler{
		Mutex:  sync.Mutex{},
		clock:  clock.VClock{},
		nodeID: uuid.New(),
		queue:  make([]uuid.UUID, 0),
	}
}
