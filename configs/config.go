package configs

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Mode string
	Port string
}

type PostgresConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	SSLmode  string
}

type Config struct {
	ServerConfig   ServerConfig
	PostgresConfig PostgresConfig
}

func InitConfig() (Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("error while reading config: %s", err.Error())
	}

	cfg := Config{
		ServerConfig: ServerConfig{
			Mode: viper.GetString("server.mode"),
			Port: viper.GetString("server.port"),
		},
		PostgresConfig: PostgresConfig{
			Host: viper.GetString("db.host"),
			Port: viper.GetString("db.port"),
			Database: viper.GetString("db.name"),
			User: viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
			SSLmode: viper.GetString("db.sslmode"),
		},
	}

	return cfg, nil
}
