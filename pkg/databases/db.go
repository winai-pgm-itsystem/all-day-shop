package databases

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/winai-pgm-itsystem/all-day-shop/config"
)

func DbConnect(cfg config.IDbConfig) *sqlx.DB {
	//Connect
	db, err := sqlx.Connect("pgx", cfg.Url())
	if err != nil {
		log.Fatalf("connect to db failed: %v", err)
	}
	db.DB.SetMaxOpenConns(cfg.MaxOpenConns())
	return db
}
