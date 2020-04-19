package config

import (
	"os"
)

// Config ties together all other application configuration types.
type Config struct {
	DbURL     string
	ServerURL string
	ProxyURL  string
}

//LoadConfig config
func LoadConfig() Config {
	var c Config
	c.DbURL = os.Getenv("DATABASE_URL")
	c.ServerURL = ":5000"
	c.ProxyURL = ":3000"
	return c
}
