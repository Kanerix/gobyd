package mutex

import (
	"sync"

	"github.com/google/uuid"
	"github.com/kanerix/gobyd/pkg/clock"
)

type LockState int

const (
	Unlocked LockState = iota + 1
	Wanted
	Locked
)

type Handler struct {
	sync.RWMutex
	NodeID    uuid.UUID
	Clock     clock.VClock
	Peers     []string
	Queue     []Request
	LockState LockState
}

func NewHandler(peers []string) *Handler {
	return &Handler{
		RWMutex:   sync.RWMutex{},
		NodeID:    uuid.New(),
		Clock:     clock.New(),
		Peers:     peers,
		Queue:     make([]Request, 0),
		LockState: Unlocked,
	}
}

func (h *Handler) TickClock() {
	h.Lock()
	defer h.Unlock()
	h.Clock.TickProcess(h.NodeID)
}
