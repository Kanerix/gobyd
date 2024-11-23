package mutex

import (
	"maps"
	"net/http"
	"slices"

	"github.com/labstack/echo/v4"
)

func (h *Handler) RequestAccess(c echo.Context) error {
	cc := c.(*Context)

	h.Lock()
	defer h.Unlock()

	if h.LockState != Unlocked {
		h.TickClock()
		h.Queue = append(h.Queue, Request{cc.requesterID, maps.Clone(h.Clock)})
		slices.SortFunc(h.Queue, SortQueue)
		return c.JSON(200, false)
	}

	return c.JSON(200, false)
}

func (h *Handler) ReleaseAccess(c echo.Context) error {
	// cc := c.(*Context)

	h.TickClock()
	if len(h.Queue) > 0 {
	}

	return nil
}

func (h *Handler) Acquire() (bool, error) {
	h.Lock()
	defer h.Unlock()

	for _, peer := range h.Peers {
		h.Clock.TickProcess(h.NodeID)

		client := http.Client{}

		req, err := http.NewRequest("GET", peer+"/api/muex/request", nil)
		if err != nil {
			return false, err
		}
		req.Header = h.Clock.IntoHeader()

		res, err := client.Do(req)
		if err != nil {
			return false, err
		}

		if res.StatusCode == 200 {
		}
	}

	return true, nil
}
