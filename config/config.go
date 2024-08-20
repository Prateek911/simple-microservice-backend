package config

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Host               string `mapstructure:"DB_HOST"`
	User               string `mapstructure:"DB_USER"`
	Password           string `mapstructure:"DB_PASSWORD"`
	Name               string `mapstructure:"DB_NAME"`
	Port               int    `mapstructure:"DB_PORT"`
	SSLMode            string `mapstructure:"DB_SSLMODE"`
	Timezone           string `mapstructure:"DB_TIMEZONE"`
	MaxPoolSize        int    `mapstructure:"MAX_POOL_SIZE"`
	IdleTimeOut        int    `mapstructure:"IDLE_TIMEOUT"`
	MaxLifetime        int    `mapstructure:"MAX_LIFETIME"`
	MaxOpenConnections int    `mapstructure:"MAX_OPEN_CONNECTIONS"`
	MaxIdleConnections int    `mapstructure:"MAX_IDLE_CONNECTIONS"`
}

type ServerOptions struct {
	MaxIdleServerConnections int `mapstructure:"MAX_IDLE_SERVER_CONNECTIONS"`
	MaxOpenServerConnections int `mapstructure:"MAX_OPEN_SERVER_CONNECTIONS"`
	DialTimeout              int `mapstructure:"DIAL_TIMEOUT"`
	Timeout                  int `mapstructure:"TIMEOUT"`
	ContextTimeOut           int `mapstructure:"CONTEXT_TIMEOUT"`
	Host                     int `mapstructure:"HOST"`
}

func GetConnectionString(cfg *DBConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode, cfg.Timezone)
}

func NewDBConfig() (*DBConfig, error) {
	_, filename, _, _ := runtime.Caller(0)
	configDir := filepath.Dir(filename)
	viper.AddConfigPath(configDir)
	viper.SetConfigName("local")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	var cfg DBConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("Error in unmarshalling DB config")
		return nil, err
	}

	return &cfg, nil
}

func NewServerConfig() (*ServerOptions, error) {
	_, filename, _, _ := runtime.Caller(0)
	configDir := filepath.Dir(filename)
	viper.AddConfigPath(configDir)
	viper.SetConfigName("local")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	var cfg ServerOptions
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("Error in unmarshalling Server config")
		return nil, err
	}

	return &cfg, nil
}
