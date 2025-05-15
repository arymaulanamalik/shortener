package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database Database
}

type Database struct {
	MysqlEnabled      bool   `envconfig:"mysql_enabled" default:"true"`
	MysqlHost         string `envconfig:"mysql_host" default:"localhost"`
	MysqlPort         int    `envconfig:"mysql_port" default:"3306"`
	MysqlDatabase     string `envconfig:"mysql_database" default:"account"`
	MysqlUser         string `envconfig:"mysql_user" default:"root"`
	MysqlPassword     string `envconfig:"mysql_password" default:"root-is-not-used"`
	MysqlMaxOpenConns int    `envconfig:"mysql_max_open_conn" default:"100"`
	MysqlMaxIdleConns int    `envconfig:"mysql_max_idle_conn" default:"10"`
}

func NewConfig() *Config {
	var conf Config
	godotenv.Load(".env")
	err := envconfig.Process("ACCOUNT", &conf)
	if err != nil {
		log.Fatalf("fail to proceed the config: %v", err)
	}
	return &conf
}
