package config

import (
	"github.com/alrusov/config"
	"github.com/alrusov/misc"
)

//----------------------------------------------------------------------------------------------------------------------------//

// Config --
type Config struct {
	Common   config.Common `toml:"common"`
	Listener HTTP          `toml:"http"`
	Others   Others        `toml:"others"`
}

// HTTP --
type HTTP struct {
	Listener config.Listener `toml:"listener"`
}

// Others  --
type Others struct {
	Option1 string `toml:"option1"`
	Option2 string `toml:"option2"`
}

//----------------------------------------------------------------------------------------------------------------------------//

// Check --
func (h *HTTP) Check(cfg *Config) error {
	var msgs []string

	if h.Listener.Timeout <= 0 {
		h.Listener.Timeout = config.ListenerDefaultTimeout
	}

	return misc.JoinedError(msgs)
}

// Check --
func (h *Others) Check(cfg *Config) error {
	var msgs []string

	if h.Option1 == "" {
		msgs = append(msgs, "Others.Option1 is empty")
	}

	if h.Option2 == "" {
		msgs = append(msgs, "Others.Option2 is empty")
	}

	return misc.JoinedError(msgs)
}

//----------------------------------------------------------------------------------------------------------------------------//

// Check --
func (cfg *Config) Check() error {
	var msgs []string

	err := cfg.Listener.Check(cfg)
	if err != nil {
		misc.AddMessage(&msgs, err.Error())
	}

	err = cfg.Others.Check(cfg)
	if err != nil {
		misc.AddMessage(&msgs, err.Error())
	}

	return misc.JoinedError(msgs)
}

//----------------------------------------------------------------------------------------------------------------------------//
