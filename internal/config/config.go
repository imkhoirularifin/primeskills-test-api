package config

import (
	"time"

	"github.com/caarlos0/env/v10"
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
}

type Database struct {
	Driver string `env:"DRIVER" envDefault:"sqlite"`
	Dsn    string `env:"DSN" envDefault:"file::memory:?cache=shared"`
}

type JwtConfig struct {
	PrivateKey string        `env:"PRIVATE_KEY,notEmpty"`
	PublicKey  string        `env:"PUBLIC_KEY,notEmpty"`
	ExpiresIn  time.Duration `env:"EXPIRES_IN" envDefault:"24h"`
	Issuer     string        `env:"ISSUER" envDefault:"localhost"`
}

type Goose struct {
	Driver       string `env:"DRIVER"`
	DbString     string `env:"DBSTRING"`
	MigrationDir string `env:"MIGRATION_DIR" envDefault:"./migrations"`
}

type Swagger struct {
	Host string `env:"HOST" envDefault:"localhost:3000"`
}

func Setup() {
	if err := env.Parse(&Cfg); err != nil {
		panic(err)
	}
}
