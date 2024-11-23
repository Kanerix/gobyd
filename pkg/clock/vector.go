package clock

import "github.com/google/uuid"

type Comparison int

const (
	EQUAL Comparison = iota + 1
	ANCESTOR
	DESCENDANT
	CONCURRENT
)

// A vector clock.
type VClock map[uuid.UUID]uint64

// Create a new vector clock.
func New() VClock {
	return make(VClock)
}

// Get the tick of a process in the vector clock (if exists).
func (vc VClock) GetProcess(nodeID uuid.UUID) (uint64, bool) {
	tick, ok := vc[nodeID]
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

// Merges the ticks of two clocks.
func (vc VClock) Merge(other VClock) {
	for nodeID, otherTick := range other {
		thisTick := vc[nodeID]
		vc[nodeID] = max(thisTick, otherTick)
	}
}
