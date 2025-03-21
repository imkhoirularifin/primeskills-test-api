package config

import (
	"time"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

var (
	Cfg Config
)

type Config struct {
	Host          string    `env:"HOST" envDefault:"0.0.0.0"`
	Port          int       `env:"PORT" envDefault:"3000"`
	IsDevelopment bool      `env:"IS_DEVELOPMENT" envDefault:"true"`
	Database      Database  `envPrefix:"DB_"`
	Jwt           JwtConfig `envPrefix:"JWT_"`
	Goose         Goose     `envPrefix:"GOOSE_"`
	Swagger       Swagger   `envPrefix:"SWAGGER_"`
	Google        Google    `envPrefix:"GOOGLE_"`
	Midtrans      Midtrans  `envPrefix:"MIDTRANS_"`
}

type Database struct {
	Driver string `env:"DRIVER" envDefault:"sqlite"`
	Dsn    string `env:"DSN" envDefault:"file::memory:?cache=shared"`
}

type JwtConfig struct {
	SecretKey string        `env:"SECRET_KEY,notEmpty"`
	ExpiresIn time.Duration `env:"EXPIRES_IN" envDefault:"24h"`
	Issuer    string        `env:"ISSUER" envDefault:"localhost"`
}

type Goose struct {
	Driver       string `env:"DRIVER"`
	DbString     string `env:"DBSTRING"`
	MigrationDir string `env:"MIGRATION_DIR" envDefault:"./migrations"`
}

type Swagger struct {
	Host string `env:"HOST" envDefault:"localhost:3000"`
}

type Google struct {
	ApplicationCredentials string `env:"APPLICATION_CREDENTIALS" envDefault:""`
}

type Midtrans struct {
	BaseUrl   string `env:"BASE_URL,notEmpty"`
	ServerKey string `env:"SERVER_KEY,notEmpty"`
}

func Setup() {
	if err := env.Parse(&Cfg); err != nil {
		panic(err)
	}
}
