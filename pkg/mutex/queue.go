package mutex

import (
	"github.com/google/uuid"
	"github.com/kanerix/gobyd/pkg/clock"
)

type Queue []Request

type Request struct {
	requesterID   uuid.UUID
	clockSnapshot clock.VClock
}

func SortQueue(i Request, j Request) int {
	return 0
}
