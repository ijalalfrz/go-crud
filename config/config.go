package config

import (
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload" // for development
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config is an app configuration.
type Config struct {
	Application struct {
		Port string
		Name string
	}
	Logger struct {
		Formatter logrus.Formatter
	}
	Mongodb struct {
		ClientOptions *options.ClientOptions
		Database      string
	}
}

// Load will load the configuration.
func Load() *Config {
	cfg := new(Config)
	cfg.logFormatter()
	cfg.app()
	cfg.mongodb()
	return cfg
}

func (cfg *Config) logFormatter() {
	formatter := &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			_, filename := path.Split(f.File)
			return funcname, filename
		},
	}

	cfg.Logger.Formatter = formatter
}

func (cfg *Config) app() {
	appName := os.Getenv("APP_NAME")
	port := os.Getenv("PORT")

	cfg.Application.Port = port
	cfg.Application.Name = appName
}

func (cfg *Config) mongodb() {
	appName := os.Getenv("APP_NAME")
	uri := os.Getenv("MONGODB_URL")
	db := os.Getenv("MONGODB_DATABASE")
	minPoolSize, _ := strconv.ParseUint(os.Getenv("MONGODB_MIN_POOL_SIZE"), 10, 64)
	maxPoolSize, _ := strconv.ParseUint(os.Getenv("MONGODB_MAX_POOL_SIZE"), 10, 64)
	maxConnIdleTime, _ := strconv.ParseInt(os.Getenv("MONGODB_MAX_IDLE_CONNECTION_TIME_MS"), 10, 64)

	opts := options.Client().
		ApplyURI(uri).
		SetAppName(appName).
		SetMinPoolSize(minPoolSize).
		SetMaxPoolSize(maxPoolSize).
		SetMaxConnIdleTime(time.Millisecond * time.Duration(maxConnIdleTime))

	cfg.Mongodb.ClientOptions = opts
	cfg.Mongodb.Database = db
}
