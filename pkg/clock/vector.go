package clock

import "github.com/google/uuid"

type VClock map[uuid.UUID]uint64

// Create a new vector clock.
func NewVClock() VClock {
	return make(VClock)
}

// Get the tick of a process in the vector clock (if exists).
func (vc VClock) GetProcess(nodeID uuid.UUID) (uint64, bool) {
	tick, ok := vc[nodeID]
	if !ok {
		vc[nodeID] = 0
		tick = 0
	}

	return tick, ok
}

// Increment the tick of a process in the vector clock.
func (vc VClock) TickProcess(nodeID uuid.UUID) {
	vc[nodeID] += 1
}

// Set the tick of a node in the vector clock.
func (vc VClock) SetTick(nodeID uuid.UUID, tick uint64) {
	vc[nodeID] = tick
}
