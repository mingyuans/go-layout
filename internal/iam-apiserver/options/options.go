package options

import (
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	"github.com/marmotedu/log"
	genericserver "github.com/mingyuans/go-layout/internal/pkg/server"
)

type Options struct {
	GenericServerRunOptions *genericserver.RunOptions             `json:"server"   mapstructure:"server"`
	InsecureServing         *genericserver.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	SecureServing           *genericserver.SecureServingOptions   `json:"secure"   mapstructure:"secure"`
	FeatureOptions          *genericserver.FeatureOptions         `json:"feature"  mapstructure:"feature"`
	Log                     *log.Options
}

func (o *Options) Flags() cliflag.NamedFlagSets {
	//这里将 options 配置为 cmd flags，允许 CMD 启动时候手动指定 value.ß
	fss := cliflag.NamedFlagSets{}
	o.GenericServerRunOptions.AddFlags(fss.FlagSet("server"))
	o.InsecureServing.AddFlags(fss.FlagSet("insecure"))
	o.SecureServing.AddFlags(fss.FlagSet("secure"))
	o.FeatureOptions.AddFlags(fss.FlagSet("feature"))
	o.Log.AddFlags(fss.FlagSet("logs"))
	return fss
}

func NewOptions() *Options {
	o := Options{
		GenericServerRunOptions: genericserver.NewRunOptions(),
		InsecureServing:         genericserver.NewInsecureServingOptions(),
		SecureServing:           genericserver.NewSecureServingOptions(),
		FeatureOptions:          genericserver.NewFeatureOptions(),
		Log:                     log.NewOptions(),
	}
	return &o
}
