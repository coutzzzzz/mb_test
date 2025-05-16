package config

import "fmt"

type Config struct {
	Service  Service
	Database Database
}

type Service struct {
	BaseURL string `mapstructure:"BASE_URL"`
	Port    string `mapstructure:"PORT"`
}

type Database struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	Name     string `mapstructure:"DB_NAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	User     string `mapstructure:"DB_USER"`
}

func (d Database) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=America/Sao_Paulo",
		d.Host,
		d.Port,
		d.User,
		d.Name,
		d.Password)
}
