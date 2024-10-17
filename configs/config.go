package configs

import (
	"strings"

	"github.com/spf13/viper"
)

var (
	current *Config
)

type SqlLite struct {
	DbPath string `mapstructure:"dbpath"`
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type Database struct {
	Driver   string    `mapstructure:"driver"`
	SqlLite  *SqlLite  `mapstructure:"sqlite"`
	Postgres *Postgres `mapstructure:"postgres"`
}

type Config struct {
	Database *Database `mapstructure:"database"`
}

func Init() (*Config, error) {
	viper.SetConfigFile("./configs/config.yml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()
	viper.WatchConfig()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	current = &config
	return &config, nil
}

func Get() *Config {
	return current
}

func Set(config *Config) {
	current = config
}
