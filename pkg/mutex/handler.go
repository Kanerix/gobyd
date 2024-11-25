package mutex

import (
	"sync"

	"github.com/google/uuid"
	"github.com/kanerix/gobyd/pkg/clock"
)

type (
	LockState int

	MutexHandler struct {
		sync.RWMutex
		NodeID    uuid.UUID
		Clock     clock.VClock
		Peers     []string
		LockState LockState
	}
)

const (
	Unlocked LockState = iota + 1
	Wanted
	Locked
)

func NewMutexHandler(peers []string) *MutexHandler {
	return &MutexHandler{
		RWMutex:   sync.RWMutex{},
		NodeID:    uuid.New(),
		Clock:     clock.New(),
		Peers:     peers,
		LockState: Unlocked,
	}
}

func (h *MutexHandler) TickClock() {
	h.Lock()
	defer h.Unlock()
	h.Clock.TickProcess(h.NodeID)
}
