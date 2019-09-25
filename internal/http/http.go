package http

import (
	"net/http"
	"sync/atomic"

	"github.com/alrusov/stdhttp"

	"github.com/alrusov/service-template/internal/config"
)

//----------------------------------------------------------------------------------------------------------------------------//

// HTTP is a relay for HTTP influxdb writes
type HTTP struct {
	cfg *config.Config
	h   *stdhttp.HTTP
}

var (
	extraInfo struct {
		Counter int64 `toml:"counter"`
	}
)

//----------------------------------------------------------------------------------------------------------------------------//

// NewHTTP --
func NewHTTP(cfg *config.Config) (*stdhttp.HTTP, error) {
	var err error

	stdhttp.SetExtraInfoFunc(
		func() interface{} {
			return &extraInfo
		},
	)

	stdhttp.AddEndpointsInfo(
		map[string]string{
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

	h := &HTTP{
		cfg: cfg,
	}

	h.h, err = stdhttp.NewListener(&cfg.Listener.Listener, h)
	if err != nil {
		return nil, err
	}

	return h.h, nil
}

//----------------------------------------------------------------------------------------------------------------------------//

// Handler --
func (h *HTTP) Handler(id uint64, path string, w http.ResponseWriter, r *http.Request) (processed bool) {
	processed = true

	switch path {
	case "/sample-url":
		atomic.AddInt64(&extraInfo.Counter, 1)
		w.Write([]byte(h.cfg.Others.Option1 + " " + h.cfg.Others.Option2))
		w.WriteHeader(http.StatusOK)

	default:
		processed = false
	}

	return
}

//----------------------------------------------------------------------------------------------------------------------------//
