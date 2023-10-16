package postgre

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"url-short/internal/config"
)

type Database struct {
	DB *sqlx.DB
}

func ConnectPSQL(cfg *config.Configuration) (*Database, error) {
	psql := new(Database)

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
		    url TEXT NOT NULL,
		    alias TEXT NOT NULL UNIQUE
		);
		
		CREATE INDEX IF NOT EXISTS idx_alias ON url_storage(alias);
	`)
	if err != nil {
		return nil, err
	}

	psql.DB = db

	return psql, nil
}

func (d *Database) CloseBD() error {
	return d.DB.Close()
}
