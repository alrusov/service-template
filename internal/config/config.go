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
	Listener config.Listener `toml:"listener"`
}

// Others -- others config
type Others struct {
	Option1 string `toml:"option1"`
	Option2 string `toml:"option2"`
}

//----------------------------------------------------------------------------------------------------------------------------//

// Check -- check http listener config
func (h *HTTP) Check(cfg *Config) error {
	msgs := misc.NewMessages()

	err := h.Listener.Check(cfg)
	if err != nil {
		msgs.Add("%s", err.Error())
	}

	return msgs.Error()
}

// Check -- check others config
func (h *Others) Check(cfg *Config) error {
	msgs := misc.NewMessages()

	if h.Option1 == "" {
		msgs.Add("others.option1 is empty")
	}

	if h.Option2 == "" {
		msgs.Add("others.option2 is empty")
	}

	return msgs.Error()
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
