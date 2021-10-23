package iam_apiserver

import (
	"github.com/marmotedu/log"
	"github.com/mingyuans/go-layout/internal/iam-apiserver/config"
	"github.com/mingyuans/go-layout/internal/iam-apiserver/options"
	"github.com/mingyuans/go-layout/internal/pkg/app"
)

const commandDesc = `The IAM API server validates and configures data
for the api objects which include users, policies, secrets, and
others. The API Server services REST operations to do the api objects management.

Find more iam-apiserver information at:
https://github.com/marmotedu/iam/blob/master/docs/guide/en-US/cmd/iam-apiserver.md`

func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp(
		//name 也用于 CMD 的 short description
		"iam-apiserver",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithRunFunc(run(opts)),
	)
	return application
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		//Init log settings
		log.Init(opts.Log)
		defer log.Flush()

		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}
		return Run(cfg)
	}
}
