package config

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"ssib.al/ssib-al-back/ent"
	"ssib.al/ssib-al-back/utils/logger"
)

var (
	dbClient *ent.Client
)

func GetDBClient() *ent.Client {
	return dbClient
}

func NewDBClient() {

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		Env.DBHost, Env.DBPort, Env.DBUser, Env.DBName, Env.DBPassword)

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		logger.Panic(err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		logger.Panic(err)
	}
	dbClient = client
	logger.Info("DB client initialized")
}
