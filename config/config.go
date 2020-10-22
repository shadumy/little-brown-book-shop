package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "192.168.1.107",
			Port:     9568,
			Username: "root",
			Password: "12345",
			Name:     "mainDb",
		},
	}
}
