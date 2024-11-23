package rest

import "github.com/kanerix/gobyd/pkg/mutex"

type Handler struct {
	*mutex.Handler
}

func NewHandler(peers []string) Handler {
	return Handler{
		Handler: mutex.NewHandler(peers),
	}
}
