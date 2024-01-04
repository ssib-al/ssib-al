package config

import (
	"fmt"
	"os"
	"strconv"

	"ssib.al/ssib-al-back/constants"
	"ssib.al/ssib-al-back/utils/logger"
)

var (
	Env    *EnvConfig
	DevEnv *DevEnvConfig
)

type DevEnvConfig struct {
	DomainPrefix string
}
type EnvConfig struct {
	Mode       string // MODE
	DBHost     string // POSTGRES_HOST
	DBPort     int    // POSTGRES_PORT
	DBUser     string // POSTGRES_USER
	DBName     string // POSTGRES_DB
	DBPassword string // POSTGRES_PASSWORD
}

func LoadStrEnv(key string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		logger.Panic(fmt.Errorf("env %s is not set", key))
	}
	return
}

func LoadIntEnv(key string) (value int) {
	strValue := os.Getenv(key)
	value, err := strconv.Atoi(strValue)
	if err != nil {
		logger.Panic(fmt.Errorf("env %s is not set or not int", key))
	}
	return
}

func NewEnvConfig() {
	Env = &EnvConfig{
		Mode:       LoadStrEnv("MODE"),
		DBHost:     LoadStrEnv("POSTGRES_HOST"),
		DBPort:     LoadIntEnv("POSTGRES_PORT"),
		DBUser:     LoadStrEnv("POSTGRES_USER"),
		DBName:     LoadStrEnv("POSTGRES_DB"),
		DBPassword: LoadStrEnv("POSTGRES_PASSWORD"),
	}
	if Env.Mode == "development" || Env.Mode == "staging" {
		logger.Info("Running in development mode")
		DevEnv = &DevEnvConfig{
			DomainPrefix: LoadStrEnv("DOMAIN_PREFIX"),
		}
		constants.InsertionPrefix(DevEnv.DomainPrefix)

	} else if Env.Mode == "production" {
		logger.Info("Running in production mode")
	} else {
		logger.Panic(fmt.Errorf("invalid mode: %s", Env.Mode))
	}
}
