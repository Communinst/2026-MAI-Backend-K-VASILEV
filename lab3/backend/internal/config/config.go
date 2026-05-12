package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type CentrifugConfig struct {
	CentrifugoURL string `env:"CENTRIFUGO_URL" env-default:"http://centrifugo:8000/api"`
	CentrifugoKey string `env:"CENTRIFUGO_API_KEY" env-default:"my_api_key"`
}

type BootConfig struct {
	ServerPort     string   `env:"SERVER_PORT" env-default:"8000"`
	ServerHost     string   `env:"SERVER_HOST" env-default:"localhost"`
	DBHost         string   `env:"DB_HOST" env-default:"localhost"`
	DBPort         string   `env:"DB_PORT" env-default:"5432"`
	DBUser         string   `env:"DB_USER" env-default:"postgres"`
	DBPassword     string   `env:"DB_PASSWORD" env-default:"postgres"`
	DBName         string   `env:"DB_NAME" env-default:"progengineering"`
	JWTSecret      string   `env:"JWT_SECRET" env-default:"secret"`
	JWTExpire      string   `env:"JWT_EXPIRE_HOURS" env-default:"24"`
	TrustedProxies []string `env:"TRUSTED_PROXIES" env-default:"127.0.0.1,::1" env-separator:","`
	TemplatesPath  string   `env:"TEMPLATES_PATH" env-default:"./templates"`
	StaticPath     string   `env:"STATIC_PATH" env-default:"./static"`
	CentriConf     CentrifugConfig
}

func LoadNewBootCfg() (*BootConfig, error) {
	var cfg BootConfig
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Failed to read boot config: %v", err)
		return nil, err
	}

	return &cfg, nil
}
