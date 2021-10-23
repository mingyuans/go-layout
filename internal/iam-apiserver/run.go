package iam_apiserver

import "github.com/mingyuans/go-layout/internal/iam-apiserver/config"

func Run(cfg *config.Config) error {
	server, err := createAPIServer(cfg)
	if err != nil {
		return err
	}
	return server.Prepare().Run()
}
