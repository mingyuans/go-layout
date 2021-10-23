// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/mingyuans/go-layout/internal/pkg/middleware"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"golang.org/x/sync/errgroup"

	"github.com/marmotedu/log"
)

// GenericAPIServer contains state for an api server.
// type GenericAPIServer gin.Engine.
type GenericAPIServer struct {
	middlewares []string

	defaultAPIs []string

	// See gin.mode
	mode string
	// SecureServingInfo holds configuration of the TLS server.
	SecureServingInfo *SecureServingOptions

	// InsecureServingInfo holds configuration of the insecure HTTP server.
	InsecureServingInfo *InsecureServingOptions

	// ShutdownTimeout is the timeout used for server shutdown. This specifies the timeout before server
	// gracefully shutdown returns.
	ShutdownTimeout time.Duration

	*gin.Engine
	enableMetrics   bool
	enableProfiling bool
	// wrapper for gin.Engine

	insecureServer, secureServer *http.Server
}

func initGenericAPIServer(s *GenericAPIServer) {
	s.Setup()
	//we placed all middlewares in pkg/middleware directory.
	s.InstallMiddlewares()
	s.InstallAPIs()
}

// InstallAPIs install generic apis.
func (s *GenericAPIServer) InstallAPIs() {
	// install metric handler
	if s.enableMetrics {
		prometheus := ginprometheus.NewPrometheus("gin")
		prometheus.Use(s.Engine)
	}

	// install pprof handler
	if s.enableProfiling {
		pprof.Register(s.Engine)
	}

	for _, name := range s.defaultAPIs {
		installAPI, ok := DefaultAPIs[name]
		if !ok {
			log.Warnf("can not find preset api: %s", name)
			continue
		}

		log.Infof("install preset api: %s", name)
		installAPI(s)
	}
}

// Setup do some setup work for gin engine.
func (s *GenericAPIServer) Setup() {
	gin.SetMode(s.mode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Infof("%-6s %-s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}
}

// InstallMiddlewares install generic middlewares.
func (s *GenericAPIServer) InstallMiddlewares() {
	// necessary middlewares
	s.Use(middleware.RequestID())
	s.Use(middleware.Context())
	// s.Use(limits.RequestSizeLimiter(10))

	// install custom middlewares
	for _, m := range s.middlewares {
		mw, ok := middleware.Middlewares[m]
		if !ok {
			log.Warnf("can not find middleware: %s", m)
			continue
		}

		log.Infof("install middleware: %s", m)
		s.Use(mw)
	}
}

/*
// preparedGenericAPIServer is a private wrapper that enforces a call of PrepareRun() before Run can be invoked.
type preparedGenericAPIServer struct {
	*GenericAPIServer
}

func (s *GenericAPIServer) PrepareRun() preparedGenericAPIServer {
	return preparedGenericAPIServer{s}
}
*/

// Run spawns the http server. It only returns when the port cannot be listened on initially.
func (s *GenericAPIServer) Run() error {
	// For scalability, use custom HTTP configuration mode here
	s.insecureServer = &http.Server{
		Addr:    s.InsecureServingInfo.Address(),
		Handler: s,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	// For scalability, use custom HTTP configuration mode here
	s.secureServer = &http.Server{
		Addr:    s.SecureServingInfo.Address(),
		Handler: s,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	var eg errgroup.Group

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	eg.Go(func() error {
		log.Infof("Start to listening the incoming requests on http address: %s", s.InsecureServingInfo.Address())

		if err := s.insecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err.Error())

			return err
		}

		log.Infof("Server on %s stopped", s.InsecureServingInfo.Address())
		return nil
	})

	eg.Go(func() error {
		key, cert := s.SecureServingInfo.ServerCert.CertKey.KeyFile, s.SecureServingInfo.ServerCert.CertKey.CertFile
		if cert == "" || key == "" || s.SecureServingInfo.BindPort == 0 {
			return nil
		}

		log.Infof("Start to listening the incoming requests on https address: %s", s.SecureServingInfo.Address())

		if err := s.secureServer.ListenAndServeTLS(cert, key); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err.Error())

			return err
		}

		log.Infof("Server on %s stopped", s.SecureServingInfo.Address())

		return nil
	})

	// Ping the server to make sure the router is working.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var isHealthzEnabled = false
	for _, name := range s.defaultAPIs {
		if name == "healthz" {
			isHealthzEnabled = true
		}
	}

	if isHealthzEnabled {
		if err := s.ping(ctx); err != nil {
			return err
		}
	}

	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

// Close graceful shutdown the api server.
func (s *GenericAPIServer) Close() {
	// The context is used to inform the server it has 10 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.secureServer.Shutdown(ctx); err != nil {
		log.Warnf("Shutdown secure server failed: %s", err.Error())
	}

	if err := s.insecureServer.Shutdown(ctx); err != nil {
		log.Warnf("Shutdown insecure server failed: %s", err.Error())
	}
}

// ping pings the http server to make sure the router is working.
func (s *GenericAPIServer) ping(ctx context.Context) error {
	url := fmt.Sprintf("http://%s/healthz", s.InsecureServingInfo.Address())
	if strings.Contains(s.InsecureServingInfo.Address(), "0.0.0.0") {
		url = fmt.Sprintf("http://127.0.0.1:%s/healthz", strings.Split(s.InsecureServingInfo.Address(), ":")[1])
	}

	for {
		// Change NewRequest to NewRequestWithContext and pass context it
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return err
		}
		// Ping the server by sending a GET request to `/healthz`.
		// nolint: gosec
		resp, err := http.DefaultClient.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			log.Info("The router has been deployed successfully.")
			//goland:noinspection GoUnhandledErrorResult
			resp.Body.Close()
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(1 * time.Second)

		select {
		case <-ctx.Done():
			log.Fatal("can not ping http server within the specified time interval.")
		default:
		}
	}
	// return fmt.Errorf("the router has no response, or it might took too long to start up")
}
