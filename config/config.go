package config

type Config struct {
	DB   *DBConfig
	Auth *AuthReqConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

type AuthReqConfig struct {
	Username string
	Password string
	Jwtkey   []byte
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
		Auth: &AuthReqConfig{
			Username: "user1",
			Password: "password",
			Jwtkey:   []byte("my_secret_key"),
		},
	}
}
