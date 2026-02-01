package db

import (
	"database/sql"
	"fmt"
	"go-crud/internal/config"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection(cfg config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	var db *sql.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Println("Successfully connected to the database!")
				return db, nil
			}
		}

		log.Printf("DB not ready... retrying in 2 seconds (Attempt %d/10): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}
	err = fmt.Errorf("Could not connect to database after 10 attempts: %v", err)
	return nil, err
}
