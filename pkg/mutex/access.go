package mutex

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

func (h *MutexHandler) RequestAccess(c echo.Context) error {
	cc := c.(*Context)

	for {
		h.Lock()

		if h.LockState != Unlocked {
			break
		}

		requesterTimestamp := h.Clock.GetProcess(cc.requesterID)
		nodeTimestamp := h.Clock.GetProcess(h.NodeID)
		if requesterTimestamp < nodeTimestamp {
			break
		}

		h.Unlock()
	}

	return c.JSON(200, true)
}

func (h *MutexHandler) ReleaseAccess(c echo.Context) error {
	// cc := c.(*Context)

	h.TickClock()
	return nil
}

func (h *MutexHandler) Acquire() {
	h.Lock()
	defer h.Unlock()

	wg := sync.WaitGroup{}
	wg.Add(len(h.Peers))

	for _, peer := range h.Peers {
		h.Clock.TickProcess(h.NodeID)
		header := h.Clock.IntoHeader()

		go func(wg *sync.WaitGroup, peer string) {
			client := http.Client{}

			req, _ := http.NewRequest("GET", peer+"/api/muex/request", nil)
			req.Header = header

			res, err := client.Do(req)
			if err != nil {
			}

			if res.StatusCode == 200 {
				wg.Done()
			}
		}(&wg, peer)
	}

	wg.Done()

	return
}
