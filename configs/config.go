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
	LogLevel string
	Server   ServerConfig
	DB       PostgresConfig
}

// Allowed logger levels & config key.
const (
	DebugLogLvl = "DEBUG"
	InfoLogLvl  = "INFO"
	ErrorLogLvl = "ERROR"
)

func InitConfig() (Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("error while reading config: %s", err.Error())
	}

	loglevel := viper.GetString("loglevel")
	if err := validate(loglevel); err != nil {
		return Config{}, fmt.Errorf("error while cheking allowed loging leveles: %w", err)
	}

	cfg := Config{
		LogLevel: loglevel,
		Server: ServerConfig{
			Mode: viper.GetString("server.mode"),
			Port: viper.GetString("server.port"),
		},
		DB: PostgresConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Database: viper.GetString("db.name"),
			User:     viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
			SSLmode:  viper.GetString("db.sslmode"),
		},
	}

	return cfg, nil
}

func validate(logLevel string) error {
	if strings.ToUpper(logLevel) != DebugLogLvl &&
		strings.ToUpper(logLevel) != ErrorLogLvl &&
		strings.ToUpper(logLevel) != InfoLogLvl {
		return fmt.Errorf("\"%v\" is not allowed logger level", logLevel)
	}

	return nil
}
