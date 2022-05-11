package echo_server

import (
	"api-gw/pkg/openapi"
	"context"
	"embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"

	log "github.com/sirupsen/logrus"

	"api-gw/controllers"
	"api-gw/pkg/config"
	"api-gw/pkg/model"
	"api-gw/pkg/utils"
	"gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/common-library.git/pkg/nexus"
)

//go:embed swagger-ui
var swaggerFS embed.FS

type EchoServer struct {
	Echo   *echo.Echo
	Config *config.Config
}

func InitEcho(stopCh chan struct{}, conf *config.Config) {
	fmt.Println("Init Echo")
	openapi.New()
	e := NewEchoServer(conf)
	e.RegisterRoutes()
	e.Start(stopCh)
}

func (s *EchoServer) StartHTTPServer() {
	if err := s.Echo.Start(":80"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error %v", err)
	}
}

func (s *EchoServer) Start(stopCh chan struct{}) {
	// Start watching URI notification
	go func() {
		log.Info("RoutesNotification")
		if err := s.RoutesNotification(stopCh); err != nil {
			s.StopServer()
			InitEcho(stopCh, s.Config)
		}
	}()

	// Start Server
	go func() {
		log.Info("Start Echo Server")
		if utils.IsServerConfigValid(s.Config) && utils.IsFileExists(s.Config.Server.CertPath) && utils.IsFileExists(s.Config.Server.KeyPath) {
			log.Infof("Server Config %v", s.Config.Server)
			log.Info("Start TLS Server")
			if err := s.Echo.StartTLS(s.Config.Server.Address, s.Config.Server.CertPath, s.Config.Server.KeyPath); err != nil && err != http.ErrServerClosed {
				log.Fatalf("TLS Server error %v", err)
			}
		} else {
			log.Info("Certificates or TLS port not configured correctly, hence starting the HTTP Server")
			s.StartHTTPServer()
		}
	}()
}

type NexusContext struct {
	echo.Context
	NexusURI string
	Codes    nexus.HTTPCodesResponse
}

func (s *EchoServer) RegisterRoutes() {
	// OpenAPI route
	s.Echo.GET("/openapi.json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, openapi.Schema)
	})

	// Swagger-UI
	fs, err := utils.GetHttpFS(swaggerFS, "swagger-ui")
	if err != nil {
		log.Warnf("Could not init swagger-ui fs: %v", err)
	}
	swaggerHandler := http.FileServer(fs)
	s.Echo.GET("/docs/*", echo.WrapHandler(http.StripPrefix("/docs/", swaggerHandler)))
}

func (s *EchoServer) RegisterRouter(restURI nexus.RestURIs) {
	urlPattern := model.ConstructEchoPathParamURL(restURI.Uri)
	for method, codes := range restURI.Methods {
		log.Infof("Registered Router Path %s Method %s\n", urlPattern, method)
		switch method {
		case http.MethodGet:
			s.Echo.GET(urlPattern, getHandler, func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
					nc := &NexusContext{
						Context:  c,
						NexusURI: restURI.Uri,
						Codes:    codes,
					}
					return next(nc)
				}
			})
		case http.MethodPut:
			s.Echo.PUT(urlPattern, putHandler, func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
					nc := &NexusContext{
						Context:  c,
						NexusURI: restURI.Uri,
						Codes:    codes,
					}
					return next(nc)
				}
			})
		case http.MethodDelete:
			s.Echo.DELETE(urlPattern, deleteHandler, func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
					nc := &NexusContext{
						Context:  c,
						NexusURI: restURI.Uri,
						Codes:    codes,
					}
					return next(nc)
				}
			})
		}
	}
}

func (s *EchoServer) RoutesNotification(stopCh chan struct{}) error {
	for {
		select {
		case <-stopCh:
			return fmt.Errorf("stop signal received")
		case restURIs := <-controllers.GlobalRestURIChan:
			log.Println("Route notification received...")
			for _, v := range restURIs {
				s.RegisterRouter(v)
				openapi.AddPath(v)
			}
		}
	}
}

func (s *EchoServer) StopServer() {
	if err := s.Echo.Shutdown(context.Background()); err != nil {
		log.Fatalf("Shutdown signal received")
	} else {
		log.Println("Server exiting")
	}
}

func NewEchoServer(conf *config.Config) *EchoServer {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	return &EchoServer{
		// create a new echo_server instance
		Echo:   e,
		Config: conf,
	}
}
