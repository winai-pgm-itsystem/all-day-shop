package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func LoadConfig(path string) IConfig {

	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("load dotenv failed: %v", err)

	}

	return &config{
		app: &app{
			host: envMap["APP_HOST"],
			port: func() int {
				p, err := strconv.Atoi(envMap["APP_PORT"])
				if err != nil {
					log.Fatalf("failed to convert  port: %v", err)
				}
				return p
			}(),
			name:    envMap["APP_NAME"],
			version: envMap["APP_VERSION"],
			readTimeOut: time.Duration(func() int {
				r, err := strconv.Atoi(envMap["APP_READ_TIMEOUT"])
				if err != nil {
					log.Fatalf("failed to convert read time out: %v", err)
				}
				return r
			}()),
		},
		db:  &db{},
		jwt: &jwt{},
	}
}

type IConfig interface {
	App() IAppConfig
	Db() IDbConfig
	Jwt() IJwtConfig
}

type config struct {
	app *app
	db  *db
	jwt *jwt
}

type IAppConfig interface {
}

type app struct {
	host         string
	port         int
	name         string
	version      string
	readTimeOut  time.Duration
	writeTimeOut time.Duration
	bodyLimit    int //bytes
	fileLimit    int //bytes
	gcp          string
}

func (c *config) App() IAppConfig {
	return nil
}

type IDbConfig interface {
}

type db struct {
	host           string
	port           int
	protocol       string
	username       string
	password       string
	database       string
	sslMode        string
	maxConnections int
}

func (c *config) Db() IDbConfig {
	return nil
}

type IJwtConfig interface {
}

type jwt struct {
	adminKey         string
	secretKey        string
	apiKey           string
	accessExpiresAt  int //sec
	refreshExpiresAt int //sec
}

func (c *config) Jwt() IJwtConfig {
	return nil
}
