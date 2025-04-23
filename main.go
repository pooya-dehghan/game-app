package main

import (
	"time"

	"github.com/pooya-dehghan/config"
	"github.com/pooya-dehghan/delivery/httpserver"
	"github.com/pooya-dehghan/repository/mysql"
	"github.com/pooya-dehghan/service/authservice"
	"github.com/pooya-dehghan/service/userservice"
)

const (
	SIGNED_KEY             = "jwt_key"
	AccessTokenExpiration  = time.Hour * 24
	RefreshTokenExpiration = time.Hour * 1
	AccessTokenSubject     = "ac"
	RefreshTokenSubject    = "rf"
)

func main() {

	cfg := config.Config{
		HTTPServer: config.HTTPServerConfig{
			Port: 8080,
		},
		Auth: authservice.Config{
			SignKey:                []byte(SIGNED_KEY),
			AccessTokenExpiration:  AccessTokenExpiration,
			RefreshTokenExpiration: RefreshTokenExpiration,
			AccessTokenSubject:     AccessTokenSubject,
			RefreshTokenSubject:    RefreshTokenSubject,
		},
		Mysql: mysql.Config{
			Username: "root",
			Password: "kashmar552",
			Host:     "localhost",
			Port:     3308,
			DBName:   "game_app",
		},
	}

	authService, userService := setupServices(cfg)

	server := httpserver.New(cfg, userService, authService)

	server.Serve()
}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service) {
	authSvc := authservice.NewService(cfg.Auth)

	mysqlRepp := mysql.New(cfg.Mysql)

	userSvc := userservice.New(mysqlRepp, cfg.Auth.SignKey)

	return authSvc, userSvc
}
