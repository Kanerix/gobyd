package muex

import "github.com/google/uuid"

type Handler struct {
	pid uuid.UUID
}

func NewHandler() Handler {
	return Handler{uuid.New()}
}
