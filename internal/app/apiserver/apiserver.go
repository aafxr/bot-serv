package apiserver

import (
	"database/sql"

	"github.com/aafxr/chat-bot/internal/app/store/sqlstore"
	_ "github.com/go-sql-driver/mysql"
)

func Start(config *Config) error {
	db, err := newDB(config.BindAddr)
	if err != nil {
		return err
	}

	defer db.Close()
	sqlstore.New(db)
	return nil
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
