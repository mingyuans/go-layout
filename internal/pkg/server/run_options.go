package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

// RunOptions contains the options while running a generic api server.
type RunOptions struct {
	Mode        string   `json:"mode"        mapstructure:"mode"`
	DefaultAPIs []string `json:"defaultapi"  mapstructure:"defaultapis"`
	Middlewares []string `json:"middlewares" mapstructure:"middlewares"`
}

// NewRunOptions creates a new RunOptions object with default parameters.
func NewRunOptions() *RunOptions {
	return &RunOptions{
		Mode:        gin.ReleaseMode,
		Middlewares: []string{},
		DefaultAPIs: []string{
			"version",
			"healthz",
			"whoami",
		},
	}
}

// Validate checks validation of ServerRunOptions.
func (s *RunOptions) Validate() []error {
	var errors []error = nil
	return errors
}

// AddFlags adds flags for a specific APIServer to the specified FlagSet.
func (s *RunOptions) AddFlags(fs *pflag.FlagSet) {
	// Note: the weird ""+ in below lines seems to be the only way to get gofmt to
	// arrange these text blocks sensibly. Grrr.
	fs.StringVar(&s.Mode, "server.mode", s.Mode, ""+
		"Start the server in a specified server mode. Supported server mode: debug, test, release.")

	fs.StringSliceVar(&s.DefaultAPIs, "server.defaultapis", s.DefaultAPIs, ""+
		"Enable some preset apis, like healthz, version,whoami...,comma separated. If this list is empty default apis will be used.")

	fs.StringSliceVar(&s.Middlewares, "server.middlewares", s.Middlewares, ""+
		"List of allowed middlewares for server, comma separated. If this list is empty default middlewares will be used.")
}
