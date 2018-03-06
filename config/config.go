package config

type Config struct {
	DBAddr             string
	DBName             string
	DBUser             string
	DBPassword         string
}

var CFG *Config