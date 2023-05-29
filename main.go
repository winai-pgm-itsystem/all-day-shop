package main

import (
	"os"

	"github.com/winai-pgm-itsystem/all-day-shop/config"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/servers"
	"github.com/winai-pgm-itsystem/all-day-shop/pkg/databases"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())

	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	servers.NewServer(cfg, db).Start()

}
