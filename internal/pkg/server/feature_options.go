package server

import (
	"github.com/spf13/pflag"
)

// FeatureOptions contains configuration items related to API server features.
type FeatureOptions struct {
	EnableProfiling bool `json:"profiling"      mapstructure:"profiling"`
	EnableMetrics   bool `json:"enable-metrics" mapstructure:"enable-metrics"`
}

// NewFeatureOptions creates a FeatureOptions object with default parameters.
func NewFeatureOptions() *FeatureOptions {
	return &FeatureOptions{
		EnableMetrics:   true,
		EnableProfiling: true,
	}
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (o *FeatureOptions) Validate() []error {
	return []error{}
}

// AddFlags adds flags related to features for a specific api server to the
// specified FlagSet.
func (o *FeatureOptions) AddFlags(fs *pflag.FlagSet) {
	if fs == nil {
		return
	}

	fs.BoolVar(&o.EnableProfiling, "feature.profiling", o.EnableProfiling,
		"Enable profiling via web interface host:port/debug/pprof/")

	fs.BoolVar(&o.EnableMetrics, "feature.enable-metrics", o.EnableMetrics,
		"Enables metrics on the apiserver at /metrics")
}
