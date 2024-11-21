package clock

import "github.com/google/uuid"

type VClock map[uuid.UUID]uint64

// Create a new vector clock.
func NewVClock() VClock {
	return make(VClock)
}

// Get the tick of a process in the vector clock (if exists).
func (vc VClock) GetProcess(pid uuid.UUID) (uint64, bool) {
	tick, ok := vc[pid]
	return tick, ok
}

// Increment the tick of a process in the vector clock
func (vc VClock) TickProcess(pid uuid.UUID) {
	vc[pid] += 1
}
