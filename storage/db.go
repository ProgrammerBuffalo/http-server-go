package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

type DBCredential struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

func DBInit() *sql.DB {
	log.Println("Start DB initialization...")

	var yamlErr error
	yamlFile, yamlErr := os.ReadFile("../db.yaml")

	if yamlErr != nil {
		log.Fatalf("Cannot read yaml file: %s", yamlErr)
	}

	var dbCred DBCredential

	if err := yaml.Unmarshal(yamlFile, &dbCred); err != nil {
		log.Fatalf("Cannot unmarshal yaml file: %s", err)
	}

	connectionString := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable",
		dbCred.Host, dbCred.Port, dbCred.User, dbCred.Password, dbCred.Name)

	var dataSource *sql.DB
	dataSource, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatalf("Database open connection err: %s", err)
	}

	if err := dataSource.Ping(); err != nil {
		log.Fatalf("Cannot ping to database: %s", err)
	}

	log.Println("End DB initialization...")
	return dataSource
}

func DBClose(ds *sql.DB) {
	log.Println("Start DB close...")
	if err := ds.Close(); err != nil {
		log.Fatalf("Cannot close database connection")
	}
	log.Println("End DB close...")
}
