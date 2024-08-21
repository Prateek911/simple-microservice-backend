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

type EnvOptions struct {
	Environment string `mapstructure:"SERVICE_VERSION"`
	Version     string `mapstructure:"NAME"`
}

func GetConnectionString(cfg *DBConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode, cfg.Timezone)
}

func loadConfig(configName string, cfg interface{}) error {
	_, filename, _, _ := runtime.Caller(1)
	configDir := filepath.Dir(filename)

	viper.AddConfigPath(configDir)
	viper.SetConfigName(configName)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("error unmarshalling config: %w", err)
	}

	return nil
}

func NewDBConfig() (*DBConfig, error) {
	var cfg DBConfig
	if err := loadConfig(".env", &cfg); err != nil {
		log.Fatalf("Failed to load DB config: %v", err)
		return nil, err
	}
	return &cfg, nil
}

func NewEnvConfig() (*EnvOptions, error) {
	var cfg EnvOptions
	if err := loadConfig(".env", &cfg); err != nil {
		log.Fatalf("Failed to load Environment config: %v", err)
		return nil, err
	}
	return &cfg, nil
}

func NewServerConfig() (*ServerOptions, error) {
	var cfg ServerOptions
	if err := loadConfig(".env", &cfg); err != nil {
		log.Fatalf("Failed to load Server config: %v", err)
		return nil, err
	}
	return &cfg, nil
}
