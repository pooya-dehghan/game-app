package config

import (
	"github.com/pooya-dehghan/repository/mysql"
	"github.com/pooya-dehghan/service/authservice"
)

type HTTPServerConfig struct {
	Port int `mapstructure:"port"`
}

type Config struct {
	HTTPServer HTTPServerConfig
	Auth       authservice.Config
	Mysql      mysql.Config
}
