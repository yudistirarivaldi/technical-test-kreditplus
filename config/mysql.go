package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection(config DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
