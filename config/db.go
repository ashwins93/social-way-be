package config

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func getConfig() *PostgresConfig {
	dbUrl, ok := os.LookupEnv("DATABASE_URL")

	validUrl, err := url.Parse(dbUrl)

	if err != nil || !ok {
		log.Fatal("Invalid db url, shutting down")
	}

	c := &PostgresConfig{}

	c.Host = validUrl.Hostname()
	c.Port = validUrl.Port()
	c.User = validUrl.User.Username()
	c.Password, _ = validUrl.User.Password()
	c.Database = validUrl.Path[1:]

	return c
}

func GetConnection() *gorm.DB {
	c := getConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", c.Host, c.Port, c.User, c.Password, c.Database)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to open db connection")
	}

	return db
}
