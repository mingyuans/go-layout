package http_options

import (
	"fmt"
	"github.com/spf13/pflag"
)

// WXClientOption Option 定义的三个要素, 第一步，先定义 Option 属性
type WXClientOption struct {
	Address string `json:"address" mapstructure:"address"`
}

// NewWXClientOption 第二步，定义 New 方法
func NewWXClientOption() *WXClientOption {
	return &WXClientOption{
		Address: "http://127.0.0.1:8081",
	}
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
// 第三步，定义 Validate
func (p *WXClientOption) Validate() []error {
	var errors []error

	if len(p.Address) == 0 {
		errors = append(
			errors,
			fmt.Errorf("--wechat.address %s must not be empty",
				p.Address,
			),
		)
	}
	return errors
}

// AddFlags adds flags related to feature for a specific api server to the
// specified FlagSet.
// 第四步，定义命令提示
func (p *WXClientOption) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&p.Address, "wx.address", p.Address, ""+
		"The wechat server address. like 127.0.0.1:8081")
}

// 第五步，将 Validate 注册到 options/options.go 底下! r
