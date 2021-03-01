package http

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/alrusov/log"
	"github.com/alrusov/misc"
	"github.com/alrusov/stdhttp"

	"github.com/alrusov/service-template/internal/config"
)

//----------------------------------------------------------------------------------------------------------------------------//

// HTTP -- listener struct
type HTTP struct {
	cfg *config.Config
	h   *stdhttp.HTTP
}

var (
	myLog = log.NewFacility("custom")

	extraInfo struct {
		Counter    int64     `json:"counter"`
		ServerTime time.Time `json:"server-time"`
	}
)

//----------------------------------------------------------------------------------------------------------------------------//

// NewHTTP -- listener initializtion
func NewHTTP(cfg *config.Config) (*stdhttp.HTTP, error) {
	var err error

	h := &HTTP{
		cfg: cfg,
	}

	h.h, err = stdhttp.NewListener(&cfg.Listener.Listener, h)
	if err != nil {
		return nil, err
	}

	h.h.SetExtraInfoFunc(
		func() interface{} {
			extraInfo.ServerTime = misc.NowUTC()
			return &extraInfo
		},
	)

	h.h.AddEndpointsInfo(
		misc.StringMap{
			"/sample-url": "Sample endpoint",
		},
	)

	h.h.SetRootItemsFunc(
		func(prefix string) []string {
			return []string{
				fmt.Sprintf(`<a href="%s/sample-url" target="sample">Sample endpoint</a>`, prefix),
				fmt.Sprintf(`<a href="%s/test.html" target="sample">File</a>`, prefix),
				`<a href="https://google.com/" target="sample">google.com</a>`,
			}
		},
	)

	return h.h, nil
}

//----------------------------------------------------------------------------------------------------------------------------//

// Handler -- custom http endpoints handler
func (h *HTTP) Handler(id uint64, prefix string, path string, w http.ResponseWriter, r *http.Request) (processed bool) {
	processed = true

	myLog.MessageWithSource(log.DEBUG, path, "Processing")

	switch path {
	case "/sample-url":
		n := atomic.AddInt64(&extraInfo.Counter, 1)
		d := []byte(fmt.Sprintf("%s %s\nCounter = %d\n", h.cfg.Others.Option1, h.cfg.Others.Option2, n))
		stdhttp.WriteReply(w, r, http.StatusOK, stdhttp.ContentTypeText, nil, d)

	default:
		processed = false
	}

	return
}

//----------------------------------------------------------------------------------------------------------------------------//
