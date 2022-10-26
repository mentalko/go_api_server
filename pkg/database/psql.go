package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type ConnectionInfo struct {
	Host     string
	Port     int
	Username string
	DBName   string
	Password string
}

func NewPostgresConnection(info ConnectionInfo) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		info.Username, info.Password, info.Host, info.Port, info.DBName))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
