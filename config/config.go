package config

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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
	file, err := os.Open("/go/src/go-with-compose/.env")

	m := make(map[string]string)
	param := []string{"DB_HOST", "DB_PASSWORD", "DB_NAME", "DB_HOST_PORT", "AUTH_USER", "AUTH_PASSWORD", "AUTH_KEY"}

	if err == nil {
		scanner := bufio.NewScanner(file)

		scanner.Split(bufio.ScanLines)
		var text []string

		for scanner.Scan() {
			text = append(text, scanner.Text())
		}

		defer file.Close()

		for _, eachLn := range text {
			// println(strings.Split(eachLn, "=")[1])
			paramChk := strings.Split(eachLn, "=")
			for _, p := range param {
				if paramChk[0] == p {
					m[p] = paramChk[1]
				}
			}
		}

	} else {
		m["DB_HOST"] = "192.168.1.107"
		m["DB_PASSWORD"] = "12345"
		m["DB_NAME"] = "mainDb"
		m["DB_HOST_PORT"] = "9568"
		m["AUTH_USER"] = "user"
		m["AUTH_PASSWORD"] = "password"
		m["AUTH_KEY"] = "my_secret_key"
	}

	dbPort, err := strconv.Atoi(m["DB_HOST_PORT"])
	if err != nil {
		dbPort = 9568
	}

	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     m["DB_HOST"],
			Port:     dbPort,
			Username: "root",
			Password: m["DB_PASSWORD"],
			Name:     m["DB_NAME"],
		},
		Auth: &AuthReqConfig{
			Username: m["AUTH_USER"],
			Password: m["AUTH_PASSWORD"],
			Jwtkey:   []byte(m["AUTH_KEY"]),
		},
	}
}
