package config

import (
	"github.com/alrusov/config"
	"github.com/alrusov/misc"
)

//----------------------------------------------------------------------------------------------------------------------------//

// Config -- application config
type Config struct {
	Common   config.Common `toml:"common"`
	Listener HTTP          `toml:"http"`
	Others   Others        `toml:"others"`
}

// HTTP -- http listener config
type HTTP struct {
	Listener      config.Listener `toml:"listener"`
}

// Others -- others config
type Others struct {
	Option1 string `toml:"option1"`
	Option2 string `toml:"option2"`
}

//----------------------------------------------------------------------------------------------------------------------------//

// Check -- check http listener config
func (h *HTTP) Check(cfg *Config) error {
	var msgs []string

	err := h.Listener.Check(cfg)
	if err != nil {
		misc.AddMessage(&msgs, err.Error())
	}

	return misc.JoinedError(msgs)
}

// Check -- check others config
func (h *Others) Check(cfg *Config) error {
	var msgs []string

	if h.Option1 == "" {
		misc.AddMessage(&msgs, "others.option1 is empty")
	}

	if h.Option2 == "" {
		misc.AddMessage(&msgs, "others.option2 is empty")
	}

	return misc.JoinedError(msgs)
}

//----------------------------------------------------------------------------------------------------------------------------//

// Check -- check application config
func (cfg *Config) Check() (err error) {
	return config.Check(
		cfg,
		[]interface{}{
			&cfg.Common,
			&cfg.Listener,
			&cfg.Others,
		},
	)
}

//----------------------------------------------------------------------------------------------------------------------------//
