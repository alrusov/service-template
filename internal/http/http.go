package http

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/alrusov/loadavg"
	"github.com/alrusov/log"
	"github.com/alrusov/misc"
	"github.com/alrusov/stdhttp"

	"github.com/alrusov/service-template/internal/config"
)

//----------------------------------------------------------------------------------------------------------------------------//

// HTTP -- listener struct
type HTTP struct {
	cfg        *config.Config
	h          *stdhttp.HTTP
	laRequests *loadavg.LoadAvg
}

var (
	extraInfo struct {
		Counter         int64   `json:"counter"`
		RequestsTotal   int64   `json:"requestsTotal"`
		RequestsLoadAvg float64 `json:"requestsLoadAvg"`
	}
)

//----------------------------------------------------------------------------------------------------------------------------//

// NewHTTP -- listener initializtion
func NewHTTP(cfg *config.Config) (*stdhttp.HTTP, error) {
	var err error

	h := &HTTP{
		cfg: cfg,
	}

	la, err := loadavg.Init(cfg.Listener.LoadAvgPeriod)
	if err != nil {
		log.Message(log.INFO, `RequestsLoadAvg: %s`, err.Error())
	} else {
		h.laRequests = la
	}

	stdhttp.SetExtraInfoFunc(
		func() interface{} {
			extraInfo.RequestsLoadAvg = h.laRequests.Value()
			return &extraInfo
		},
	)

	stdhttp.AddEndpointsInfo(
		misc.StringMap{
			"/sample-url": "Sample endpoint",
		},
	)

	stdhttp.SetRootItemsFunc(
		func() []string {
			return []string{
				`<a href="/sample-url" target="sample">Sample URL</a>`,
				`<a href="https://google.com/" target="sample">google.com</a>`,
			}
		},
	)

	h.h, err = stdhttp.NewListener(&cfg.Listener.Listener, h)
	if err != nil {
		return nil, err
	}

	return h.h, nil
}

//----------------------------------------------------------------------------------------------------------------------------//

// Handler -- custom http endpoints handler
func (h *HTTP) Handler(id uint64, path string, w http.ResponseWriter, r *http.Request) (processed bool) {
	processed = true

	atomic.AddInt64(&extraInfo.RequestsTotal, 1)
	h.laRequests.Add(1)

	switch path {
	case "/sample-url":
		n := atomic.AddInt64(&extraInfo.Counter, 1)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%s %s\nCounter = %d\n", h.cfg.Others.Option1, h.cfg.Others.Option2, n)))

	default:
		processed = false
	}

	return
}

//----------------------------------------------------------------------------------------------------------------------------//
