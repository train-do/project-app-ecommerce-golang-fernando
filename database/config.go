package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/train-do/project-app-ecommerce-golang-fernando/utils"
)

func InitDB(config utils.Configuration) (*sql.DB, error) {
	// connStr := "user=postgres dbname=ecommerce-db sslmode=disable password=superUser host=localhost"
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s", config.DB.Username, config.DB.Password, config.DB.Name, config.DB.Host)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)
	return db, err
}
