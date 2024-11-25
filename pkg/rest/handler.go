package rest

import "github.com/kanerix/gobyd/pkg/mutex"

type RestHandler struct {
	*mutex.MutexHandler
}

func NewRestHandler(peers []string) *RestHandler {
	return &RestHandler{
		MutexHandler: mutex.NewMutexHandler(peers),
	}
}
