package config

import (
	"database/sql"
	"fmt"
	"log"
)

func (s Database) CreateDatabaseIfNotExists(driver string) error {
	connectionInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", s.Username, s.Password, s.Host, s.Port, s.DefaultDatabase)
	log.Printf("connectionInfo: %s", connectionInfo)
	db, openErr := sql.Open(driver, connectionInfo)
	if openErr != nil {
		return openErr
	}
	defer db.Close()
	if !s.IsExistsDatabase(db, s.UseDatabase) {
		err := s.CreateDatabase(db, s.UseDatabase)
		if err != nil {
			log.Printf("CreateDatabase Error: %s failed to create: %s", err, s.UseDatabase)
			return err
		} else {
			log.Printf("Database '%s' created successfully.\n", s.UseDatabase)
		}
	}
	return nil
}

func (s Database) IsExistsDatabase(db *sql.DB, database string) bool {
	var dbExists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", database).Scan(&dbExists)
	if err != nil {
		log.Printf("db error: %s, %s does not exists", err, database)
		return false
	}
	return dbExists
}

func (s Database) CreateDatabase(db *sql.DB, database string) error {
	_, err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", database))
	if err != nil {
		return err
	}
	return nil
}
