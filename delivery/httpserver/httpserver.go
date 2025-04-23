package httpserver

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pooya-dehghan/config"
	"github.com/pooya-dehghan/service/authservice"
	"github.com/pooya-dehghan/service/userservice"
)

type Config struct {
	Port int
}

type Server struct {
	config      config.Config
	userService userservice.Service
	authService authservice.Service
}

func New(config config.Config, userService userservice.Service, authService authservice.Service) Server {
	return Server{
		config:      config,
		userService: userService,
		authService: authService,
	}
}

func (s Server) Serve() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/health-check", s.healthCheck)

	userGroup := e.Group("/users")

	userGroup.POST("/users/register", s.userRegister)
	userGroup.POST("/users/login", s.userLogin)
	userGroup.POST("/users/profile", s.userProfile)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))
}
