package configs

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	JwtSecret    string
	JwtExp       string
	ConnPostgres string
}

func GetConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	var (
		port         = os.Getenv("PORT")
		jwtSecret    = os.Getenv("JWT_SECRET")
		jwtExp       = os.Getenv("JWT_EXP")
		connPostgres = os.Getenv("CONN_POSTGRES")
	)

	if port == "" || jwtSecret == "" || jwtExp == "" || connPostgres == "" {
		return nil, errors.New("error of parsing .env file")
	}

	return &Config{
		Port:         ":" + port,
		JwtSecret:    jwtSecret,
		JwtExp:       jwtExp,
		ConnPostgres: connPostgres,
	}, nil
}
