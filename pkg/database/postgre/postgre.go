package postgre

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"url-short/internal/config"
)

func ConnectDB(cfg *config.Configuration) (*sqlx.DB, error) {
	db, err := sqlx.Connect(
		cfg.Database.Driver,
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.DbName, cfg.Database.Password))
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.Database.DbIdleTimeout)

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS url_storage (
		    id SERIAL PRIMARY KEY,
		    url VARCHAR,
		    alias VARCHAR(50)
		)
	`)
	if err != nil {
		return nil, err
	}

	return db, err
}
