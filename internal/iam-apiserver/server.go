package iam_apiserver

import (
	"github.com/marmotedu/iam/pkg/shutdown"
	"github.com/marmotedu/iam/pkg/shutdown/shutdownmanagers/posixsignal"
	"github.com/marmotedu/log"
	"github.com/mingyuans/go-layout/internal/iam-apiserver/config"
	genericserver "github.com/mingyuans/go-layout/internal/pkg/server"
)

type apiServer struct {
	shutdown      *shutdown.GracefulShutdown
	restAPIServer *genericserver.GenericAPIServer
}

type preparedAPIServer struct {
	*apiServer
}

func createAPIServer(cfg *config.Config) (*apiServer, error) {
	gs := shutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	genericConfig := buildGenericConfig(cfg)
	cmptConfig := buildComponentConfig(cfg)

	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	err = cmptConfig.complete().New()
	if err != nil {
		return nil, err
	}

	server := &apiServer{
		shutdown:      gs,
		restAPIServer: genericServer,
	}
	return server, nil
}

func buildGenericConfig(cfg *config.Config) *genericserver.Config {
	genericConfig := genericserver.NewConfig()
	if cfg.SecureServing != nil {
		genericConfig.SecureServing = cfg.SecureServing
	}
	if cfg.InsecureServing != nil {
		genericConfig.InsecureServing = cfg.InsecureServing
	}
	if cfg.FeatureOptions != nil {
		genericConfig.FeatureInfo = cfg.FeatureOptions
	}
	if cfg.GenericServerRunOptions != nil {
		genericConfig.RunInfo = cfg.GenericServerRunOptions
	}
	return genericConfig
}

func buildComponentConfig(cfg *config.Config) *componentConfig {
	componentConfig := newComponentConfig()
	return componentConfig
}

func (s *apiServer) Prepare() preparedAPIServer {
	installControllers(s.restAPIServer.Engine)

	//初始化其他的 components
	//s.initRedis()

	s.shutdown.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		//关闭 components
		//mysqlStore, _ := mysql.GetMySQLFactoryOr(nil)
		//if mysqlStore != nil {
		//	return mysqlStore.Close()
		//}

		// close services
		s.restAPIServer.Close()

		return nil
	}))

	return preparedAPIServer{s}
}

func (s preparedAPIServer) Run() error {
	// start shutdown managers
	if err := s.shutdown.Start(); err != nil {
		log.Fatalf("start shutdown manager failed: %s", err.Error())
	}

	return s.restAPIServer.Run()
}
