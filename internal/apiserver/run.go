package apiserver

import "github.com/mingyuans/go-layout/internal/apiserver/config"

func Run(cfg *config.Config) error {
	server, err := createAPIServer(cfg)
	if err != nil {
		return err
	}
	return server.Prepare().Run()
}
