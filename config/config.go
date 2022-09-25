package config

import (
	"sync"

	"github.com/IqbalLx/alterra-agmc/utils"
)

var lock = &sync.Mutex{}

var config *Config

type Config struct {
	Database *DatabaseConfig
	Auth     *AuthConfig
	Server   *ServerConfig
}

type DatabaseConfig struct {
	Host  string
	Port  string
	User  string
	Pwd   string
	Table string
}

type AuthConfig struct {
	JWTSecret string
}

type ServerConfig struct {
	Port string
}

func GetConfig() *Config {
	// singleton pattern
	if config == nil {
		env := utils.NewEnv(".env")

		lock.Lock()
		defer lock.Unlock()
		if config == nil {
			config = &Config{
				Database: &DatabaseConfig{
					Host:  env.Read("DATABASE_HOST"),
					Port:  env.Read("DATABASE_PORT"),
					User:  env.Read("DATABASE_USER"),
					Pwd:   env.Read("DATABASE_PWD"),
					Table: env.ReadWithDefaultVal("DATABASE_TABLE", utils.Default{Value: "agmc"}),
				},
				Auth: &AuthConfig{
					JWTSecret: env.Read("AUTH_SECRET_KEY"),
				},
				Server: &ServerConfig{
					env.ReadWithDefaultVal("PORT", utils.Default{Value: "3000"}),
				},
			}
		}
	}
	return config
}
